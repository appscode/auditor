/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_test

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	api "kubeshield.dev/auditor/apis/auditor/v1alpha1"
	"kubeshield.dev/auditor/test/e2e/framework"

	"github.com/appscode/go/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	dynamic_util "kmodules.xyz/client-go/dynamic"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

const (
	retryTimeout = 5 * time.Minute
)

var _ = Describe("Grafana Operator E2E testing", func() {
	var (
		f *framework.Invocation
	)

	var (
		getModelFromAuditRegistrationJson = func(filename string) []byte {
			By("Selecting the dashboard model")
			file, err := os.Open(filename)
			Expect(err).NotTo(HaveOccurred())

			modelData, err := ioutil.ReadAll(file)
			Expect(err).NotTo(HaveOccurred())

			return modelData
		}

		createAppBindingAndSecret = func() {
			By("Creating AppBinding " + f.AppBindingName())
			err := f.CreateAppBinding()
			Expect(err).NotTo(HaveOccurred())

			By("Creating Secret " + f.SecretName())
			err = f.CreateSecret(options.apiKey)
			Expect(err).NotTo(HaveOccurred())
		}

		waitForAuditRegistrationToGetToFinalPhase = func() {
			By("Waiting for the dashboard to get to final phase")
			Eventually(func() bool {
				dashboard, err := f.GetAuditRegistration()
				Expect(err).NotTo(HaveOccurred())

				return dashboard.Status.Phase == api.AuditRegistrationPhaseSuccess || dashboard.Status.Phase == api.AuditRegistrationPhaseFailed
			}, retryTimeout, 100*time.Millisecond).Should(BeTrue())

		}

		waitForAuditRegistrationToBeTerminated = func(dashboard *api.AuditRegistration) {
			By("Deleting dashboard " + dashboard.Name)
			err := f.DeleteAuditRegistration()
			Expect(err).NotTo(HaveOccurred())

			By("Waiting for the dashboard to be terminated")
			Eventually(func() error {
				_, err = f.GetAuditRegistration()
				return err
			}, retryTimeout, 100*time.Millisecond).ShouldNot(BeNil())
		}
	)

	BeforeEach(func() {
		f = root.Invoke()
	})

	Describe("AuditRegistration Operation", func() {
		var (
			dashboard *api.AuditRegistration
		)

		BeforeEach(func() {
			dashboard = &api.AuditRegistration{
				ObjectMeta: metav1.ObjectMeta{
					Name:      f.Name(),
					Namespace: f.Namespace(),
				},
				Spec: api.AuditRegistrationSpec{
					Grafana: &api.TargetRef{
						Name: f.AppBindingName(),
					},
					Overwrite: true,
				},
			}

		})

		Context("Successful creation of a dashboard resource", func() {
			BeforeEach(func() {
				model := getModelFromAuditRegistrationJson("dashboard-with-panels-with-mixed-yaxes.json")
				dashboard.Spec.Model = &runtime.RawExtension{
					Raw: model,
				}

				createAppBindingAndSecret()
			})

			It("should insert a dashboard into grafana database", func() {
				err := f.CreateAuditRegistration(dashboard)
				Expect(err).NotTo(HaveOccurred())

			})

			AfterEach(func() {
				By("Verifying dashboard has been succeeded")
				dashboard, err := f.GetAuditRegistration()
				Expect(err).NotTo(HaveOccurred())

				Expect(dashboard.Status.Phase).To(BeEquivalentTo(api.AuditRegistrationPhaseSuccess))

				Expect(dashboard.Status.AuditRegistration.ID).NotTo(BeNil())
				Expect(dashboard.Status.AuditRegistration.UID).NotTo(BeNil())
				Expect(dashboard.Status.AuditRegistration.Version).To(BeEquivalentTo(types.Int64P(dashboard.Status.ObservedGeneration)))
			})
		})

		Context("Unsuccessful creation of a dashboard resource", func() {
			BeforeEach(func() {
				model := getModelFromAuditRegistrationJson("dashboard-with-panels-with-mixed-yaxes.json")
				dashboard.Spec.Model = &runtime.RawExtension{
					Raw: model,
				}

				options.apiKey = "unauthorized-api-key"
				createAppBindingAndSecret()
			})

			It("should not insert a dashboard into grafana database", func() {
				err := f.CreateAuditRegistration(dashboard)
				Expect(err).NotTo(HaveOccurred())

			})

			AfterEach(func() {
				dashboard, err := f.GetAuditRegistration()
				Expect(err).NotTo(HaveOccurred())
				Expect(dashboard.Status.Phase).To(BeEquivalentTo(api.AuditRegistrationPhaseFailed))
				Expect(dashboard.Status.AuditRegistration).To(BeNil())
			})
		})

		Context("Unsuccessful creation of a dashboard resource", func() {
			BeforeEach(func() {
				By("Not providing model data")
				createAppBindingAndSecret()
			})

			It("should not insert a dashboard into grafana database", func() {
				err := f.CreateAuditRegistration(dashboard)
				Expect(err).NotTo(HaveOccurred())

			})

			AfterEach(func() {
				dashboard, err := f.GetAuditRegistration()
				Expect(err).NotTo(HaveOccurred())
				Expect(dashboard.Status.Phase).To(BeEquivalentTo(api.AuditRegistrationPhaseFailed))
				Expect(dashboard.Status.AuditRegistration).To(BeNil())
			})
		})

		JustAfterEach(func() {
			waitForAuditRegistrationToGetToFinalPhase()
		})

		AfterEach(func() {
			waitForAuditRegistrationToBeTerminated(dashboard)

			By("Deleting Secret")
			err := root.DeleteSecret()
			Expect(err).NotTo(HaveOccurred())

			By("Deleting AppBinding")
			err = root.DeleteAppBinding()
			Expect(err).NotTo(HaveOccurred())

			By("Wait until AppBinding is deleted")
			ctx := context.TODO()
			ctx, cancel := context.WithTimeout(ctx, retryTimeout)
			defer cancel()
			ri := dynamic.NewForConfigOrDie(f.RestConfig()).
				Resource(appcat.SchemeGroupVersion.WithResource(appcat.ResourceApps)).
				Namespace(f.Namespace())
			err = dynamic_util.WaitUntilDeleted(ri, ctx.Done(), f.AppBindingName())
			Expect(err).NotTo(HaveOccurred())
		})
	})
})

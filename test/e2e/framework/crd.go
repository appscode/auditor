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

package framework

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyCRD() GomegaAsyncAssertion {
	return Eventually(
		func() error {
			// Check Datasource CRD
			/*if _, err := f.extClient.AuditorV1alpha1().Datasources(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return fmt.Errorf("CRD AuditRegistration is not ready. Reason: %v", err)
			}*/

			// Check AuditRegistration CRD
			if _, err := f.extClient.AuditorV1alpha1().AuditRegistrations(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return fmt.Errorf("CRD AuditRegistration is not ready. Reason: %v", err)
			}

			// Check AuditRegistrationTemplate CRD
			/*if _, err := f.extClient.AuditorV1alpha1().AuditRegistrationTemplates(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return fmt.Errorf("CRD AuditRegistration is not ready. Reason: %v", err)
			}*/
			return nil
		},
		time.Minute*2,
		time.Second*10,
	)
}

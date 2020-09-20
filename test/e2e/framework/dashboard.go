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

	api "kubeshield.dev/auditor/apis/auditor/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
)

func (f *Framework) GetAuditRegistration() (*api.AuditRegistration, error) {
	return f.extClient.AuditorV1alpha1().AuditRegistrations(f.namespace).Get(context.TODO(), f.name, metav1.GetOptions{})
}

func (f *Framework) CreateAuditRegistration(dashboard *api.AuditRegistration) error {
	_, err := f.extClient.AuditorV1alpha1().AuditRegistrations(dashboard.Namespace).Create(context.TODO(), dashboard, metav1.CreateOptions{})
	return err
}

func (f *Framework) DeleteAuditRegistration() error {
	return f.extClient.AuditorV1alpha1().AuditRegistrations(f.namespace).Delete(context.TODO(), f.name, meta_util.DeleteInForeground())
}

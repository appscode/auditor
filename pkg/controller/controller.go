/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	api "kubeshield.dev/auditor/apis/auditor/v1alpha1"
	cs "kubeshield.dev/auditor/client/clientset/versioned"
	"kubeshield.dev/auditor/client/clientset/versioned/typed/auditor/v1alpha1/util"
	grafanainformers "kubeshield.dev/auditor/client/informers/externalversions"
	"kubeshield.dev/auditor/pkg/eventer"

	"github.com/golang/glog"
	core "k8s.io/api/core/v1"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/apiextensions"
)

type GrafanaController struct {
	config
	clientConfig *rest.Config

	kubeClient    kubernetes.Interface
	dynamicClient dynamic.Interface
	crClient      cs.Interface
	crdClient     crd_cs.Interface
	recorder      record.EventRecorder

	dynamicInformerFactory dynamicinformer.DynamicSharedInformerFactory
	crInformerFactory      grafanainformers.SharedInformerFactory
}

func (c *GrafanaController) ensureCustomResourceDefinitions() error {
	crds := []*apiextensions.CustomResourceDefinition{
		api.Dashboard{}.CustomResourceDefinition(),
	}
	return apiextensions.RegisterCRDs(c.crdClient, crds)
}

func (c *GrafanaController) Run(stopCh <-chan struct{}) {
	go c.RunInformers(stopCh)

	<-stopCh
}

func (c *GrafanaController) RunInformers(stopCh <-chan struct{}) {
	defer runtime.HandleCrash()

	glog.Info("Starting Auditor")

	c.crInformerFactory.Start(stopCh)
	for _, v := range c.crInformerFactory.WaitForCacheSync(stopCh) {
		if !v {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return
		}
	}

	<-stopCh
	glog.Info("Stopping Auditor")
}

func (c *GrafanaController) pushFailureEvent(dashboard *api.Dashboard, reason string) {
	c.recorder.Eventf(
		dashboard,
		core.EventTypeWarning,
		eventer.EventReasonFailedToStart,
		`Failed to complete operation for Dashboard: "%v". Reason: %v`,
		dashboard.Name,
		reason,
	)
	dashboard, err := util.UpdateDashboardStatus(context.TODO(), c.crClient.AuditorV1alpha1(), dashboard.ObjectMeta, func(in *api.DashboardStatus) *api.DashboardStatus {
		in.Phase = api.DashboardPhaseFailed
		in.Reason = reason
		in.Conditions = kmapi.SetCondition(in.Conditions, kmapi.Condition{
			Type:    kmapi.ConditionFailure,
			Status:  kmapi.ConditionTrue,
			Reason:  reason,
			Message: reason,
		})
		in.ObservedGeneration = dashboard.Generation
		return in
	}, metav1.UpdateOptions{})
	if err != nil {
		c.recorder.Eventf(
			dashboard,
			core.EventTypeWarning,
			eventer.EventReasonFailedToUpdate,
			err.Error(),
		)
	}
}

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
	"time"

	cs "kubeshield.dev/auditor/client/clientset/versioned"
	grafanainformers "kubeshield.dev/auditor/client/informers/externalversions"
	"kubeshield.dev/auditor/pkg/eventer"

	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"kmodules.xyz/client-go/discovery"
)

type config struct {
	DockerRegistry string
	MaxNumRequeues int
	NumThreads     int
	ResyncPeriod   time.Duration
}

type Config struct {
	config

	ClientConfig  *rest.Config
	KubeClient    kubernetes.Interface
	DynamicClient dynamic.Interface
	CRClient      cs.Interface
	CRDClient     crd_cs.Interface
}

func NewConfig(clientConfig *rest.Config) *Config {
	return &Config{
		ClientConfig: clientConfig,
	}
}

func (c *Config) New() (*AuditorController, error) {
	if err := discovery.IsDefaultSupportedVersion(c.KubeClient); err != nil {
		return nil, err
	}

	ctrl := &AuditorController{
		config:                 c.config,
		clientConfig:           c.ClientConfig,
		kubeClient:             c.KubeClient,
		dynamicClient:          c.DynamicClient,
		crClient:               c.CRClient,
		crdClient:              c.CRDClient,
		dynamicInformerFactory: dynamicinformer.NewDynamicSharedLiteInformerFactory(c.DynamicClient, c.ResyncPeriod),
		crInformerFactory:      grafanainformers.NewSharedInformerFactory(c.CRClient, c.ResyncPeriod),
		recorder:               eventer.NewEventRecorder(c.KubeClient, "auditor"),
	}

	if err := ctrl.ensureCustomResourceDefinitions(); err != nil {
		return nil, err
	}

	// For Dashboard
	ctrl.initDashboardWatcher()

	return ctrl, nil
}

/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clusterapiprovideropenstackapiv1alpha1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha1"
	clientset "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/clientset/clientset"
	internalinterfaces "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/informers/externalversions/internalinterfaces"
	apiv1alpha1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/listers/api/v1alpha1"
)

// OpenStackServerInformer provides access to a shared informer and lister for
// OpenStackServers.
type OpenStackServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() apiv1alpha1.OpenStackServerLister
}

type openStackServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOpenStackServerInformer constructs a new informer for OpenStackServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOpenStackServerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOpenStackServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOpenStackServerInformer constructs a new informer for OpenStackServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOpenStackServerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InfrastructureV1alpha1().OpenStackServers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InfrastructureV1alpha1().OpenStackServers(namespace).Watch(context.TODO(), options)
			},
		},
		&clusterapiprovideropenstackapiv1alpha1.OpenStackServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *openStackServerInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOpenStackServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *openStackServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterapiprovideropenstackapiv1alpha1.OpenStackServer{}, f.defaultInformer)
}

func (f *openStackServerInformer) Lister() apiv1alpha1.OpenStackServerLister {
	return apiv1alpha1.NewOpenStackServerLister(f.Informer().GetIndexer())
}

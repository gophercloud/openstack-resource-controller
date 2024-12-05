/*
Copyright 2024 The ORC Authors.

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
	"context"
	time "time"

	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	clientset "github.com/k-orc/openstack-resource-controller/pkg/clients/clientset/clientset"
	internalinterfaces "github.com/k-orc/openstack-resource-controller/pkg/clients/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/listers/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PortInformer provides access to a shared informer and lister for
// Ports.
type PortInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PortLister
}

type portInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPortInformer constructs a new informer for Port type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPortInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPortInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPortInformer constructs a new informer for Port type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPortInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenstackV1alpha1().Ports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenstackV1alpha1().Ports(namespace).Watch(context.TODO(), options)
			},
		},
		&apiv1alpha1.Port{},
		resyncPeriod,
		indexers,
	)
}

func (f *portInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPortInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *portInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiv1alpha1.Port{}, f.defaultInformer)
}

func (f *portInformer) Lister() v1alpha1.PortLister {
	return v1alpha1.NewPortLister(f.Informer().GetIndexer())
}

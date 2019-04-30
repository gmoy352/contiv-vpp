// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	bgpconfigv1 "github.com/contiv/vpp/plugins/crd/pkg/apis/bgpconfig/v1"
	versioned "github.com/contiv/vpp/plugins/crd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/contiv/vpp/plugins/crd/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/contiv/vpp/plugins/crd/pkg/client/listers/bgpconfig/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BgpConfigInformer provides access to a shared informer and lister for
// BgpConfigs.
type BgpConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BgpConfigLister
}

type bgpConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBgpConfigInformer constructs a new informer for BgpConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBgpConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBgpConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBgpConfigInformer constructs a new informer for BgpConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBgpConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BgpconfigV1().BgpConfigs(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BgpconfigV1().BgpConfigs(namespace).Watch(options)
			},
		},
		&bgpconfigv1.BgpConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *bgpConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBgpConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bgpConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bgpconfigv1.BgpConfig{}, f.defaultInformer)
}

func (f *bgpConfigInformer) Lister() v1.BgpConfigLister {
	return v1.NewBgpConfigLister(f.Informer().GetIndexer())
}

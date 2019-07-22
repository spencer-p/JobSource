/*
Copyright 2019 The Knative Authors

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
	time "time"

	jobsourcev1alpha1 "github.com/spencer-p/jobsource/pkg/apis/jobsource/v1alpha1"
	versioned "github.com/spencer-p/jobsource/pkg/client/clientset/versioned"
	internalinterfaces "github.com/spencer-p/jobsource/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/spencer-p/jobsource/pkg/client/listers/jobsource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// JobSourceInformer provides access to a shared informer and lister for
// JobSources.
type JobSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.JobSourceLister
}

type jobSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewJobSourceInformer constructs a new informer for JobSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewJobSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredJobSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredJobSourceInformer constructs a new informer for JobSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredJobSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JobsourceV1alpha1().JobSources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JobsourceV1alpha1().JobSources(namespace).Watch(options)
			},
		},
		&jobsourcev1alpha1.JobSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *jobSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredJobSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *jobSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jobsourcev1alpha1.JobSource{}, f.defaultInformer)
}

func (f *jobSourceInformer) Lister() v1alpha1.JobSourceLister {
	return v1alpha1.NewJobSourceLister(f.Informer().GetIndexer())
}
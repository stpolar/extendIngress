/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/extendingress-controller/pkg/apis/extendingresscontroller/v1alpha1"
)

// ExtendIngressLister helps list ExtendIngresses.
type ExtendIngressLister interface {
	// List lists all ExtendIngresses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ExtendIngress, err error)
	// ExtendIngresses returns an object that can list and get ExtendIngresses.
	ExtendIngresses(namespace string) ExtendIngressNamespaceLister
	ExtendIngressListerExpansion
}

// extendIngressLister implements the ExtendIngressLister interface.
type extendIngressLister struct {
	indexer cache.Indexer
}

// NewExtendIngressLister returns a new ExtendIngressLister.
func NewExtendIngressLister(indexer cache.Indexer) ExtendIngressLister {
	return &extendIngressLister{indexer: indexer}
}

// List lists all ExtendIngresses in the indexer.
func (s *extendIngressLister) List(selector labels.Selector) (ret []*v1alpha1.ExtendIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExtendIngress))
	})
	return ret, err
}

// ExtendIngresses returns an object that can list and get ExtendIngresses.
func (s *extendIngressLister) ExtendIngresses(namespace string) ExtendIngressNamespaceLister {
	return extendIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExtendIngressNamespaceLister helps list and get ExtendIngresses.
type ExtendIngressNamespaceLister interface {
	// List lists all ExtendIngresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ExtendIngress, err error)
	// Get retrieves the ExtendIngress from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ExtendIngress, error)
	ExtendIngressNamespaceListerExpansion
}

// extendIngressNamespaceLister implements the ExtendIngressNamespaceLister
// interface.
type extendIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExtendIngresses in the indexer for a given namespace.
func (s extendIngressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ExtendIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExtendIngress))
	})
	return ret, err
}

// Get retrieves the ExtendIngress from the indexer for a given namespace and name.
func (s extendIngressNamespaceLister) Get(name string) (*v1alpha1.ExtendIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("extendingress"), name)
	}
	return obj.(*v1alpha1.ExtendIngress), nil
}

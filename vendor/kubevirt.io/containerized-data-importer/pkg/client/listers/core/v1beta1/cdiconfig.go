/*
Copyright 2018 The CDI Authors.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// CDIConfigLister helps list CDIConfigs.
// All objects returned here must be treated as read-only.
type CDIConfigLister interface {
	// List lists all CDIConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CDIConfig, err error)
	// Get retrieves the CDIConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.CDIConfig, error)
	CDIConfigListerExpansion
}

// cDIConfigLister implements the CDIConfigLister interface.
type cDIConfigLister struct {
	indexer cache.Indexer
}

// NewCDIConfigLister returns a new CDIConfigLister.
func NewCDIConfigLister(indexer cache.Indexer) CDIConfigLister {
	return &cDIConfigLister{indexer: indexer}
}

// List lists all CDIConfigs in the indexer.
func (s *cDIConfigLister) List(selector labels.Selector) (ret []*v1beta1.CDIConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CDIConfig))
	})
	return ret, err
}

// Get retrieves the CDIConfig from the index for a given name.
func (s *cDIConfigLister) Get(name string) (*v1beta1.CDIConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("cdiconfig"), name)
	}
	return obj.(*v1beta1.CDIConfig), nil
}
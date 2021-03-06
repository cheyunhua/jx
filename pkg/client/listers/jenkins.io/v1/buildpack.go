// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BuildPackLister helps list BuildPacks.
type BuildPackLister interface {
	// List lists all BuildPacks in the indexer.
	List(selector labels.Selector) (ret []*v1.BuildPack, err error)
	// BuildPacks returns an object that can list and get BuildPacks.
	BuildPacks(namespace string) BuildPackNamespaceLister
	BuildPackListerExpansion
}

// buildPackLister implements the BuildPackLister interface.
type buildPackLister struct {
	indexer cache.Indexer
}

// NewBuildPackLister returns a new BuildPackLister.
func NewBuildPackLister(indexer cache.Indexer) BuildPackLister {
	return &buildPackLister{indexer: indexer}
}

// List lists all BuildPacks in the indexer.
func (s *buildPackLister) List(selector labels.Selector) (ret []*v1.BuildPack, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BuildPack))
	})
	return ret, err
}

// BuildPacks returns an object that can list and get BuildPacks.
func (s *buildPackLister) BuildPacks(namespace string) BuildPackNamespaceLister {
	return buildPackNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BuildPackNamespaceLister helps list and get BuildPacks.
type BuildPackNamespaceLister interface {
	// List lists all BuildPacks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.BuildPack, err error)
	// Get retrieves the BuildPack from the indexer for a given namespace and name.
	Get(name string) (*v1.BuildPack, error)
	BuildPackNamespaceListerExpansion
}

// buildPackNamespaceLister implements the BuildPackNamespaceLister
// interface.
type buildPackNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BuildPacks in the indexer for a given namespace.
func (s buildPackNamespaceLister) List(selector labels.Selector) (ret []*v1.BuildPack, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BuildPack))
	})
	return ret, err
}

// Get retrieves the BuildPack from the indexer for a given namespace and name.
func (s buildPackNamespaceLister) Get(name string) (*v1.BuildPack, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("buildpack"), name)
	}
	return obj.(*v1.BuildPack), nil
}

package postgres

import (
	"k8s.io/apimachinery/pkg/runtime"
)

type APIObjectVersioner struct{}

func (a *APIObjectVersioner) UpdateObject(obj runtime.Object, resourceVersion uint64) error {
	panic("implement me")
}

func (a *APIObjectVersioner) UpdateList(obj runtime.Object, resourceVersion uint64, continueValue string, remainingItemCount *int64) error {
	panic("implement me")
}

func (a *APIObjectVersioner) PrepareObjectForStorage(obj runtime.Object) error {
	panic("implement me")
}

func (a *APIObjectVersioner) ObjectResourceVersion(obj runtime.Object) (uint64, error) {
	panic("implement me")
}

func (a *APIObjectVersioner) ParseResourceVersion(resourceVersion string) (uint64, error) {
	panic("implement me")
}

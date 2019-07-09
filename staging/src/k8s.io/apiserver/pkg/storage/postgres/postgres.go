package postgres

import (
	"context"
	"database/sql"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
)

type store struct {
	db        *sql.DB
	versioner storage.Versioner
}

func New(db *sql.DB) storage.Interface {
	return newStore(db)
}

func newStore(db *sql.DB) *store {
	versioner := &APIObjectVersioner{}
	return &store{
		db:        db,
		versioner: versioner,
	}
}

func (s *store) Versioner() storage.Versioner {
	return s.versioner
}

func (store) Create(ctx context.Context, key string, obj, out runtime.Object, ttl uint64) error {
	panic("implement me")
}

func (store) Delete(ctx context.Context, key string, out runtime.Object, preconditions *storage.Preconditions, validateDeletion storage.ValidateObjectFunc) error {
	panic("implement me")
}

func (store) Watch(ctx context.Context, key string, resourceVersion string, p storage.SelectionPredicate) (watch.Interface, error) {
	panic("implement me")
}

func (store) WatchList(ctx context.Context, key string, resourceVersion string, p storage.SelectionPredicate) (watch.Interface, error) {
	panic("implement me")
}

func (store) Get(ctx context.Context, key string, resourceVersion string, objPtr runtime.Object, ignoreNotFound bool) error {
	panic("implement me")
}

func (store) GetToList(ctx context.Context, key string, resourceVersion string, p storage.SelectionPredicate, listObj runtime.Object) error {
	panic("implement me")
}

func (store) List(ctx context.Context, key string, resourceVersion string, p storage.SelectionPredicate, listObj runtime.Object) error {
	panic("implement me")
}

func (store) GuaranteedUpdate(
	ctx context.Context, key string, ptrToType runtime.Object, ignoreNotFound bool,
	precondtions *storage.Preconditions, tryUpdate storage.UpdateFunc, suggestion ...runtime.Object) error {
	panic("implement me")
}

func (store) Count(key string) (int64, error) {
	panic("implement me")
}

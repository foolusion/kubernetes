package factory

import (
	"database/sql"
	_ "github.com/lib/pq"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/postgres"
	"k8s.io/apiserver/pkg/storage/storagebackend"
)

func newPostgresClient(c storagebackend.Config) (*sql.DB, DestroyFunc, error) {
	connStr := "user=kubernetes dbname=kubernetes sslmode=verify-full sslcert=" + c.Transport.CertFile + " sslkey=" + c.Transport.KeyFile + " sslrootcert=" + c.Transport.CAFile
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func newPostgresStorage(c storagebackend.Config) (storage.Interface, DestroyFunc, error) {
	db, destroyFunc, err := newPostgresClient(c)
	if err != nil {
		return nil, nil, err
	}
	return postgres.New(db), destroyFunc, nil
}

func newPostgresHealthCheck(c storagebackend.Config) (func() error, error) {
	return func() error { return nil }, nil
}

package oss

import (
	"io"
	"os"
	"time"
)

// StorageInterface define common API to operate storage
type StorageInterface interface {
	GetURL(path string) (string, error)
	GetStream(path string) (io.ReadCloser, error)
	Get(path string) (*os.File, error)
	Put(path string, reader io.Reader) (*Object, error)
	Delete(path string) error
	List(path string) ([]*Object, error)
	GetEndpoint() string
}

// Object content object
type Object struct {
	Path             string
	Name             string
	LastModified     *time.Time
	StorageInterface StorageInterface
}

// Get retrieve object's content
func (object Object) Get() (*os.File, error) {
	return object.StorageInterface.Get(object.Path)
}

// GetStream retrieve object's content as stream
func (object Object) GetStream() (io.ReadCloser, error) {
	return object.StorageInterface.GetStream(object.Path)
}

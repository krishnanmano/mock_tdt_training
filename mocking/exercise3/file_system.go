package exercise3

import (
	"os"
)

// FileSystem is an interface for interacting with the file system.
type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte) error
}

// RealFileSystem is the implementation of FileSystem that uses the real file system.
type RealFileSystem struct {
	// Some real file system implementation here...
}

// ReadFile reads data from the specified file.
func (fs *RealFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// WriteFile writes data to the specified file.
func (fs *RealFileSystem) WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, os.ModePerm)
}

// NewRealFileSystem creates a new instance of RealFileSystem.
func NewRealFileSystem() *RealFileSystem {
	return &RealFileSystem{}
}

func WriteToFile(file FileSystem, filename string, data []byte) error {
	return file.WriteFile(filename, data)
}

func ReadToFile(file FileSystem, filename string) ([]byte, error) {
	return file.ReadFile(filename)
}

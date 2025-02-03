package files

import "path/filepath"

// This module is implementation of Storage interface for
// storing files locally

type Local struct {
	maxFileSize int //max bytes of file size
	basePath    string
}

// NewLocal creates a new local filesystem with given basePath
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil
}

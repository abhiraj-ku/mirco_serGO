package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// This module is implementation of Storage interface for
// storing files locally

type Local struct {
	maxFileSize int //max bytes of file size
	basePath    string
}

// NewLocal creates a new local filesystem with given basePath
// basePath is the base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil
}

// save the contents of writer to the given path

func (l *Local) Save(path string, contents io.Reader) error {
	// full path for the file
	fp := l.fullPath(path)

	// get the dir & make sure it exists
	dir := filepath.Dir(fp)
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("unable to create directory : %w", err)
	}

	// if the file exists just delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("Unable to delete the file %w", err)
		}
	} else if !os.IsExist(err) {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	// create a new file (upar hamne dir banaya h)
	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file %w", err)
	}

	defer f.Close()

	// wrire the io.Reader contents to this newly created file
	// io.Reader --> f (newly created file above)

	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to the dest file:%w", err)
	}

	return nil

}

// Get the file at given path and return a Reader

func (l *Local) Get(path string) (*os.File, error) {
	//get the full path of the file

	fp := l.fullPath(path)

	//open the file
	f, err := os.Open(fp)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open the file %w", err)
	}

	return f, nil
}

// Fullpath function

func (l *Local) fullPath(path string) string {
	// append the given path to base path

	return filepath.Join(l.basePath, path)
}

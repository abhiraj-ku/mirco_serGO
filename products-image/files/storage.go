package files

import "io"

type Storage interface {
	save(path string, file io.Reader) error
}

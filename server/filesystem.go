package server

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrIsDirectory is returned when it's a directory.
	ErrIsDirectory = errors.New("is a directory")
	// ErrUnsafePath is returned when given path is unsafe.
	ErrUnsafePath = errors.New("unsafe or invalid path specified")
)

type Filesystem struct {
	root string
}

// NewFilesystem returns a new Filesystem.
func NewFilesystem(root string) *Filesystem {
	return &Filesystem{
		root: root,
	}
}

// Sanitize returns a normalized path, preventing path traversal.
func (fs *Filesystem) Sanitize(p string) (string, error) {
	p = filepath.Clean(filepath.Join(fs.root, strings.TrimPrefix(p, fs.root)))

	r, err := filepath.EvalSymlinks(p)
	if err != nil {
		return "", errors.Wrap(err, "could not evaluate symlinks")
	}
	// If the directory from EvalSymlinks begins with filesystem root directory, return it.
	// Otherwise, we will just return an error to prevent further action.
	if strings.HasPrefix(strings.TrimSuffix(r, "/"), strings.TrimSuffix(fs.root, "/")) {
		return r, nil
	}
	return "", ErrUnsafePath
}

// File returns a reader, and information for a file.
func (fs *Filesystem) File(p string) (*os.File, os.FileInfo, error) {
	sanitized, err := fs.Sanitize(p)
	if err != nil {
		return nil, nil, err
	}

	st, err := os.Stat(sanitized)
	if err != nil {
		return nil, nil, err
	}
	if st.IsDir() {
		return nil, nil, ErrIsDirectory
	}

	f, err := os.Open(sanitized)
	if err != nil {
		return nil, nil, err
	}
	return f, st, nil
}

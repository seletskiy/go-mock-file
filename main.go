// Package mockfile provides a way for testing packages that rely upon using
// os.File interface.
package mockfile

import (
	"io"
	"os"
	"time"
)

type file struct {
	name string
	buf  []byte
	ptr  int64
}

// New creates new mock file, which can be used as os.File.
func New(name string) *file {
	return &file{
		name: name,
		buf:  []byte{},
		ptr:  0,
	}
}

func (m *file) Close() error {
	return nil
}

func (m *file) Read(p []byte) (n int, err error) {
	n, err = m.ReadAt(p, m.ptr)
	m.ptr += int64(n)

	return n, err
}

func (m *file) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case os.SEEK_SET:
		m.ptr = offset
	case os.SEEK_END:
		m.ptr = int64(len(m.buf)) + offset
	case os.SEEK_CUR:
		m.ptr = m.ptr + offset
	}

	return m.ptr, nil
}

func (m *file) Stat() (os.FileInfo, error) {
	return fileInfo{int64(len(m.buf))}, nil
}

func (m *file) ReadAt(p []byte, off int64) (n int, err error) {
	if n = copy(p, m.buf[off:]); n == 0 {
		return n, io.EOF
	} else {
		return n, nil
	}
}

type fileInfo struct {
	size int64
}

func (m fileInfo) Name() string {
	return ""
}

func (m fileInfo) Size() int64 {
	return m.size
}

func (m fileInfo) Mode() os.FileMode {
	return os.FileMode(0)
}

func (m fileInfo) ModTime() time.Time {
	return time.Time{}
}

func (m fileInfo) IsDir() bool {
	return false
}

func (m fileInfo) Sys() interface{} {
	return nil
}

func (m *file) Write(p []byte) (n int, err error) {
	m.buf = append(m.buf, p...)

	return len(p), nil
}

func (m *file) Truncate(size int64) error {
	if size > int64(len(m.buf)) {
		size = int64(len(m.buf))
	}

	m.buf = m.buf[:size-1]

	return nil
}

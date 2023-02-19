package model

import (
	"io"
	"time"
)

type ObjUnwrap interface {
	Unwrap() Obj
}

type Obj interface {
	GetSize() int64
	GetName() string
	ModTime() time.Time
	IsDir() bool

	// The internal information of the driver.
	// If you want to use it, please understand what it means
	GetID() string
	GetPath() string
}

type FileStreamer interface {
	io.ReadCloser
	Obj
	GetMimetype() string
	SetReadCloser(io.ReadCloser)
	NeedStore() bool
	GetReadCloser() io.ReadCloser
	GetOld() Obj
}

package webdav

import (
	"errors"
	"net/http"
)

type Handler struct {
	Prefix string

	LockSystem LockSystem

	Logger func(*http.Request, error)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := http.StatusBadRequest, errUnsupportedMethod
	brw := newBufferedResponseWriter()
	useBufferedWriter := true
	if h.LockSystem == nil {
		status, err = http.StatusInternalServerError, errNoLockSystem
	} else {
		switch r.Method {
		case "OPTIONS":
			status, err = h.han
		}
	}
}

var (
	errNoLockSystem            = errors.New("webdav: no lock system")
	errUnsupportedMethod       = errors.New("webdav: unsupported method")
)
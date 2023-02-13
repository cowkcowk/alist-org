package webdav

import "strconv"

type LockSystem interface {
	Confirm() (release func(), err error)

	Create() (token string, err error)

	Refresh() (error)

	Unlock() error
}

type memLS struct {
	byName  map[string]*memLSNode
	byToken map[string]*memLSNode
	gen     uint64
}

func (m *memLS) nextToken() string {
	m.gen++
	return strconv.FormatUint(m.gen, 10)
}

func (m *memLS) Confirm() (func(), error) {
	return func() {

	}, nil
}

func (m *memLS) Create() (string, error) {
	return "", nil
}

func (m *memLS) Refresh() (error) {
	return nil
}

func (m *memLS) Unlock() error {
	return nil
}

type memLSNode struct {
}

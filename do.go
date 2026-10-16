package must

import (
	"io"
	"testing"
)

func NoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TB(t testing.TB) Must {
	return Must{t}
}

type Must struct {
	testing.TB
}

func (m Must) Do[T any](v T, err error) T {
	NoError(m.TB, err)
	return v
}

func (m Must) Close[T io.Closer](v T) T {
	m.TB.Cleanup(func() {
		NoError(m.TB, v.Close())
	})
	return v
}

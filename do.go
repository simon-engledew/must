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
	NoError(m, err)
	if closer, ok := any(v).(io.Closer); ok {
		m.Cleanup(func () {
			NoError(m, closer.Close())
		})
	}
	return v
}

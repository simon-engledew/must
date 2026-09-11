package must_test

import (
	"net/url"
	"testing"

	"github.com/simon-engledew/must"
)

func TestDo(t *testing.T) {
	v := must.TB(t).Do(url.Parse("http://google.com"))
	if v.Host != "google.com" {
		t.Fatal("wrong host")
	}
}

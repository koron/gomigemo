package inflator

import (
	"testing"
)

func TestSuffix(t *testing.T) {
	s := Suffix("foo", "bar", "baz")
	c := s.Inflate("qux-")
	if <-c != "qux-foo" {
		t.Error("Suffix didn't return \"qux-foo\"")
	}
	if <-c != "qux-bar" {
		t.Error("Suffix didn't return \"qux-bar\"")
	}
	if <-c != "qux-baz" {
		t.Error("Suffix didn't return \"qux-baz\"")
	}
	if _, ok := <-c; ok {
		t.Error("Suffix returned unexpected")
	}
}

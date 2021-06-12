package inflator

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	p := Prefix("foo", "bar", "baz")
	c := p.Inflate("-qux")
	if <-c != "foo-qux" {
		t.Error("Prefix didn't return \"foo-qux\"")
	}
	if <-c != "bar-qux" {
		t.Error("Prefix didn't return \"bar-qux\"")
	}
	if <-c != "baz-qux" {
		t.Error("Prefix didn't return \"baz-qux\"")
	}
	if _, ok := <-c; ok {
		t.Error("Prefix returned unexpected")
	}
}

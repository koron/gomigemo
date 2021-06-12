package inflator

import (
	"testing"
)

func TestEcho(t *testing.T) {
	e := Echo()

	c1 := e.Inflate("foo")
	if <-c1 != "foo" {
		t.Error("Echo didn't return \"foo\"")
	}
	if _, ok := <-c1; ok {
		t.Error("Echo returned others of \"foo\"")
	}

	c2 := e.Inflate("bar")
	if <-c2 != "bar" {
		t.Error("Echo didn't return \"bar\"")
	}
	if _, ok := <-c2; ok {
		t.Error("Echo returned others of \"bar\"")
	}
}

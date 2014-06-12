package migemo

import (
	"testing"
)

func assertNotNull(t *testing.T, o interface{}) {
	if o == nil {
		t.Errorf("detect nul unexpectedly")
	}
}

func TestMigemoNew(t *testing.T) {
	g := New("./dict")
	assertNotNull(t, g)
}

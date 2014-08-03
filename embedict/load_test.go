package embedict

import (
	"testing"
)

func TestLoad(t *testing.T) {
	d, err := Load()
	if err != nil {
		t.Fatal("failed to load embedded dict", err)
	}
	if d == nil {
		t.Fatal("embedict.Load returns nil")
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Load()
	}
}

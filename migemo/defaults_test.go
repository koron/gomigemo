package migemo

import (
	"testing"
)

func BenchmarkLoadDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = LoadDefault()
	}
}

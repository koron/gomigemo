package migemo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplitClauses(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want []string
	}{
		{"abc", []string{"abc"}},
		{"abcDef", []string{"abc", "Def"}},
		{"abcDEFghi", []string{"abc", "DEF", "ghi"}},

		{"aaa0", []string{"aaa0"}},
		{"AAA0", []string{"AAA", "0"}},
	} {
		got, err := splitClauses(tc.in)
		if err != nil {
			t.Errorf("unexpected error at #%d (%+v): %s", i, tc, err)
		}
		if d := cmp.Diff(tc.want, got); d != "" {
			t.Errorf("output mismatch: -want +got\n%s", d)
		}
	}
}

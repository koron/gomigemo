package skk

import (
	"testing"
)

func TestNewLoader(t *testing.T) {
	l := NewLoader("http://openlab.jp/skk/dic/SKK-JISYO.L.gz")
	if l == nil {
		t.Error("NewLoader() returns nil")
	}
}

func parseEntry(t *testing.T, s string) (entry Entry) {
	e, err := parseJisyoEntry(s)
	if err != nil {
		t.Errorf("parseJisyoEntry(\"%s\") causes an error: %s", s, err)
	}
	entry = e.(Entry)
	return
}

func assertEntry(t *testing.T, expected, actually Entry) {
	if actually.Key != expected.Key {
		t.Errorf("Entry.Key are mismatched: expected=%s actually=%s",
			expected.Key, actually.Key)
	}
	if len(actually.Values) != len(expected.Values) {
		t.Errorf(
			"length of Entry.Values are mismatched: expected=%d actually=%d",
			len(expected.Values), len(actually.Values))
	}
	for i, s := range actually.Values {
		if s != expected.Values[i] {
			t.Errorf(
				"Entry.Values[%d] are mismatched: expected=%s actually=%s",
				i, expected.Values[i], s)
		}
	}
}

func checkParseEntry(t *testing.T, expected Entry, s string) {
	assertEntry(t, expected, parseEntry(t, s))
}

func TestParseJisyoEntry(t *testing.T) {
	_, err := parseJisyoEntry("foo")
	if err == nil {
		t.Error("\"foo\" MUST be error")
	}
	checkParseEntry(t,
		Entry{"foo", []string{"bar"}},
		"foo /bar/")
	checkParseEntry(t,
		Entry{"foo", []string{"bar", "baz"}},
		"foo /bar/baz/")
	checkParseEntry(t,
		Entry{"foo", []string{"bar", "qux"}},
		"foo /bar;baz/qux;quux/")
}

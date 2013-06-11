package migemo

import (
	"testing"
)

func assertEquals(t *testing.T, exp string, act string) {
	if exp != act {
		t.Errorf("expected=%s actually=%s", exp, act)
	}
}

func assertConvert(t *testing.T, c *SConv, exp string, in string) {
	out, err := c.Convert(in)
	if err != nil {
		t.Error("failed to convert", err)
	}
	assertEquals(t, exp, out)
}

func TestSConvEmpty(t *testing.T) {
	c := NewSConv()
	assertConvert(t, c, "", "")
	assertConvert(t, c, "foo", "foo")
	assertConvert(t, c, "bar", "bar")
}

func TestSConvSimple(t *testing.T) {
	c := NewSConv()
	c.Add("a", "A", "")
	c.Add("b", "B", "")
	assertConvert(t, c, "A", "a")
	assertConvert(t, c, "B", "b")
	assertConvert(t, c, "c", "c")
	assertConvert(t, c, "AAAAABBBBBccccc", "aaaaabbbbbccccc")
}

func TestSConvTiny(t *testing.T) {
	c := NewSConv()
	c.Add("aa", "A", "a")
	c.Add("ab", "B", "")
	assertConvert(t, c, "B", "ab")
	assertConvert(t, c, "Aa", "aa")
	assertConvert(t, c, "AB", "aab")
	assertConvert(t, c, "Bc", "abc")
	assertConvert(t, c, "Ba", "aba")
}

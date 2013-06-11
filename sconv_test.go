package migemo

import (
    "testing"
)

func assertEquals(t *testing.T, exp string, act string) {
    if (exp != act) {
        t.Errorf("expected=%s actually=%s", exp, act)
    }
}

func checkConvert(t *testing.T, c *SConv, exp string, in string) {
    out, err := c.Convert(in)
    if err != nil {
        t.Error("failed to convert", err)
    }
    assertEquals(t, exp, out)
}

func TestSConvEmpty(t *testing.T) {
    c := NewSConv()
    checkConvert(t, c, "", "")
    checkConvert(t, c, "foo", "foo")
    checkConvert(t, c, "bar", "bar")
}

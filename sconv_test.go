package migemo

import (
    "testing"
)

func assertEquals(t *testing.T, exp string, act string) {
    if (exp != act) {
        t.Errorf("expected=%s actually=%s", exp, act)
    }
}

func TestSConvEmpty(t *testing.T) {
    c := NewSConv()
    assertEquals(t, "", c.Convert(""))
    assertEquals(t, "foo", c.Convert("foo"))
    assertEquals(t, "bar", c.Convert("bar"))
}

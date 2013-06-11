package migemo

import (
	"testing"
)

func TestMatcher(t *testing.T) {
	m := NewMatcher()
	assertEquals(t, "", m.Pattern())
	m.Add("a")
	assertEquals(t, "a", m.Pattern())
	m.Add("b")
	assertEquals(t, "[ab]", m.Pattern())
	m.Add("A")
	assertEquals(t, "[Aab]", m.Pattern())
	m.Add("AB")
	assertEquals(t, "[Aab]", m.Pattern())
	m.Add("CD")
	assertEquals(t, "(?:[Aab]|C\\s*D)", m.Pattern())
	m.Add("CE")
	assertEquals(t, "(?:[Aab]|C\\s*[DE])", m.Pattern())
	m.Add("C")
	assertEquals(t, "[ACab]", m.Pattern())
}

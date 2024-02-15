package migemo

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// mcDict means multiple clause (renbun) dictionary.
type mcDict struct {
	dict Dict
}

// MultiClauses creates a new Dict which supports multiple clauses (renbun)
// conversion based on given Dict.
func MultiClauses(dict Dict) Dict {
	return mcDict{dict: dict}
}

// Matcher returns a new Matcher object which generated from query.

func (d mcDict) Matcher(query string) (Matcher, error) {
	clauses, err := splitClauses(query)
	if err != nil {
		return nil, err
	}
	if len(clauses) == 0 {
		return nil, errors.New("no clauses in query")
	}
	mm := make(mcMatcher, len(clauses))
	for i, c := range clauses {
		m, err := d.dict.Matcher(strings.ToLower(c))
		if err != nil {
			return nil, fmt.Errorf("clause #%d %q failed to create a matcher: %w", i, c, err)
		}
		mm[i] = m
	}
	return mm, nil
}

type mcMatcher []Matcher

func (mm mcMatcher) Match(s string) (chan Match, error) {
	// TODO:
	return nil, errors.New("Match method of multi-clause is not implemented yet")
}

func (mm mcMatcher) Pattern() (string, error) {
	var bb bytes.Buffer
	o := mm[0].GetOptions()
	for i, m := range mm {
		p, err := m.Pattern()
		if err != nil {
			return "", fmt.Errorf("clause #%d failed to generate pattern: %w", i, err)
		}
		// Consider MatcherOptions when concatenate.
		if bb.Len() > 0 {
			bb.WriteString(o.OpWSpaces)
		}
		bb.WriteString(p)
	}
	return bb.String(), nil
}

func (mm mcMatcher) SetOptions(o MatcherOptions) {
	for _, m := range mm {
		m.SetOptions(o)
	}
}

func (mm mcMatcher) GetOptions() MatcherOptions {
	return mm[0].GetOptions()
}

// splitClauses separates a string into clauses.  The break between clauses is
// usually an uppercase letter.  Clauses that begin with multiple capital
// letters are separated by non-capital letters.
func splitClauses(query string) ([]string, error) {
	a := make([]string, 0, 8)
	mode := 0
	cstart := -1
	for i, ch := range query {
		if mode == 0 {
			cstart = i
			if unicode.IsUpper(ch) {
				mode = 2
				continue
			}
			mode = 1
			continue
		}
		// start with lower char.
		if mode == 1 {
			if unicode.IsUpper(ch) {
				mode = 2
				a = append(a, query[cstart:i])
				cstart = i
				continue
			}
			continue
		}
		// start with upper char, process 2nd char
		if mode == 2 {
			if unicode.IsUpper(ch) {
				mode = 3
				continue
			}
			mode = 1
			continue
		}
		// start with two upper chars
		if mode == 3 {
			if !unicode.IsUpper(ch) {
				mode = 1
				a = append(a, query[cstart:i])
				cstart = i
				continue
			}
			continue
		}
	}
	if cstart < len(query) {
		a = append(a, query[cstart:])
	}
	return a, nil
}

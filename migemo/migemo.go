package migemo

import (
	"regexp"
)

// Dict provides an interface of dictionary for Migemo.
type Dict interface {
	Matcher(string) (Matcher, error)
}

// Matcher defines Migemo matcher interface, which have an expanded migemo tree
// and provides matching operations.
type Matcher interface {
	// Match matches with string.
	Match(string) (chan Match, error)
	// Pattern provides a regexp pattern of this match.
	Pattern() (string, error)
	// SetOptions changes matcher's options.
	SetOptions(MatcherOptions)
	// GetOptions retrieves matcher's options.
	GetOptions() MatcherOptions
}

// MatcherOptions defines options for migemo matcher (generation of regexp).
type MatcherOptions struct {
	OpOr                  string
	OpGroupIn, OpGroupOut string
	OpClassIn, OpClassOut string
	OpWSpaces             string
	//MetaChars string
}

// Match is positional information of a match.
type Match struct {
	Start, End int
}

// Load loads a dict from path (file system).
func Load(path string) (Dict, error) {
	return LoadAssets(&PathAssets{root: path})
}

// LoadAssets loads a dict from Assets.
func LoadAssets(assets Assets) (Dict, error) {
	d := &dict{assets: assets}
	err := d.load()
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Compile compiles a regexp which expanded from string s with Migemo.
func Compile(d Dict, s string) (*regexp.Regexp, error) {
	m, err := d.Matcher(s)
	if err != nil {
		return nil, err
	}
	return NewRegexp(m)
}

// NewRegexp generates a regexp from matcher.
func NewRegexp(m Matcher) (*regexp.Regexp, error) {
	p, err := m.Pattern()
	if err != nil {
		return nil, err
	}
	return regexp.Compile(p)
}

// Pattern generates a regexp patter string from string s with Migemo.
func Pattern(d Dict, s string) (string, error) {
	m, err := d.Matcher(s)
	if err != nil {
		return "", err
	}
	p, err := m.Pattern()
	if err != nil {
		return "", err
	}
	return p, nil
}

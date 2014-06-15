package migemo

var defaultMatcherOptions = MatcherOptions{
	OpOr:       "|",
	OpGroupIn:  "[",
	OpGroupOut: "]",
	OpClassIn:  "(?:",
	OpClassOut: ")",
	OpWSpaces:  "\\s+",
	// TODO: list up regexp meta characters.
	MetaChars:  "",
}

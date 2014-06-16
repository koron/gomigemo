package migemo

var defaultMatcherOptions = MatcherOptions{
	OpOr:       "|",
	OpGroupIn:  "(?:",
	OpGroupOut: ")",
	OpClassIn:  "[",
	OpClassOut: "]",
	OpWSpaces:  "\\s+",
	// FIXME: Support MetaChars customization in future.
	//MetaChars:  "",
}

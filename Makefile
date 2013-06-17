include $(GOROOT)/src/Make.inc

TARG=migemo
GOFILES=\
	matcher.go\
	migemo.go\
	scon.go\
	stacked_reader.go\
	ternary_trie.go\

include $(GOROOT)/src/Make.pkg

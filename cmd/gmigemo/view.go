package main

type View interface {
	GetQuery() (string, error)
	PutPattern(string) error
}

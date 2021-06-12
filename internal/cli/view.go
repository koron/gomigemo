package cli

// View provides interface for gmigemo's input and ouput.
type View interface {
	GetQuery() (string, error)
	PutPattern(string) error
}

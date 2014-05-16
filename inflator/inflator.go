package inflator

type Inflatable interface {
	Inflate(s string) Inflator
}

type Inflator interface {
	NextString() (string, bool)
}

package inflator

type echogen struct {
}

type echo struct {
	seed string
	valid bool
}

var echogenSingle = echogen{}

func Echo() Inflatable {
	return &echogenSingle
}

func (e *echogen) Inflate(s string) Inflator {
	return &echo{s, true}
}

func (e *echo) NextString() (string, bool) {
	if e.valid {
		e.valid = false
		return e.seed, true
	} else {
		return "", false
	}
}

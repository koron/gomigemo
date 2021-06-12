package inflator

type echo struct {
}

// Echo provides an Inflatable which returns a string as is.
func Echo() Inflatable {
	return &echo{}
}

func (e *echo) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		c <- s
	})
}

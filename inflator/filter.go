package inflator

type filter struct {
	check func(string) bool
}

// Filter provides a filter Inflatable.
func Filter(check func(string) bool) Inflatable {
	return &filter{check}
}

func (f *filter) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		if f.check(s) {
			c <- s
		}
	})
}

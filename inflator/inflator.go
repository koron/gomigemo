package inflator

// Inflatable inflates a string with channel.
type Inflatable interface {
	Inflate(s string) <-chan string
}

// Start starts inflation with function.
func Start(f func(chan<- string)) <-chan string {
	c := make(chan string, 1)
	go func() {
		defer close(c)
		f(c)
	}()
	return c
}

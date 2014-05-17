package inflator

type Inflatable interface {
	Inflate(s string) <-chan string
}

func Start(f func(chan<- string)) <-chan string {
	c := make(chan string, 1)
	go func() {
		defer close(c)
		f(c)
	}()
	return c
}

////////////////////////////////////////////////////////////////////////////
// Dispatch

type dispatcher struct {
	inflatables []Inflatable
}

func Dispatch(v ...Inflatable) Inflatable {
	return &dispatcher{v}
}

func (d *dispatcher) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		for _, n := range d.inflatables {
			for t := range n.Inflate(s) {
				c <- t
			}
		}
	})
}

////////////////////////////////////////////////////////////////////////////
// Filter

type filter struct {
	check func(string) bool
}

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

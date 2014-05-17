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
// Join

type joiner struct {
	first, second Inflatable
}

func Join(first, second Inflatable) Inflatable {
	return &joiner{first, second}
}

func (j *joiner) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		for t := range j.first.Inflate(s) {
			for u := range j.second.Inflate(t) {
				c <- u
			}
		}
	})
}

////////////////////////////////////////////////////////////////////////////
// Echo

type echo struct {
}

func Echo() Inflatable {
	return &echo{}
}

func (e *echo) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		c <- s
	})
}

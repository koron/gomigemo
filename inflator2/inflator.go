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

////////////////////////////////////////////////////////////////////////////
// Prefix

type prefixer struct {
	prefixes []string
}

func Prefix(prefixes ...string) Inflatable {
	return &prefixer{prefixes}
}

func (p *prefixer) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		for _, t := range p.prefixes {
			c <- t + s
		}
	})
}

////////////////////////////////////////////////////////////////////////////
// Suffix

type suffixer struct {
	suffixes []string
}

func Suffix(suffixes ...string) Inflatable {
	return &suffixer{suffixes}
}

func (p *suffixer) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		for _, t := range p.suffixes {
			c <- s + t
		}
	})
}

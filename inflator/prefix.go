package inflator

type prefixer struct {
	prefixes []string
}

// Prefix provides Inflatable with prefixes.
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

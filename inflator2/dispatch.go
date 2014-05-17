package inflator

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

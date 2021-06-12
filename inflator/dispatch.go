package inflator

type dispatcher struct {
	inflatables []Inflatable
}

// Dispatch is an Inflatable which dispatch a string to Inflatables.
func Dispatch(first Inflatable, others ...Inflatable) Inflatable {
	inflatables := make([]Inflatable, len(others)+1)
	inflatables[0] = first
	for i, v := range others {
		inflatables[i+1] = v
	}
	return &dispatcher{inflatables}
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

// DispatchEcho is an Inflatable which combined Echo and Dispatch.
func DispatchEcho(inflatables ...Inflatable) Inflatable {
	return Dispatch(Echo(), inflatables...)
}

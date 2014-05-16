package inflator

func Dispatch(v ...Inflatable) Inflatable {
	d := &dispatchable{v}
	return d
}

type dispatchable struct {
	inflatables []Inflatable
}

type dinflator struct {
	seed        string
	inflatables []Inflatable
	nextIndex   int
	current     Inflator
}

func (d *dispatchable) Inflate(s string) Inflator {
	return &dinflator{s, d.inflatables, 0, nil}
}

func (d *dinflator) NextString() (string, bool) {
	for {
		if d.current == nil {
			if d.nextIndex >= len(d.inflatables) {
				return "", false
			}
			d.current = d.inflatables[d.nextIndex].Inflate(d.seed)
		}

		if s, has := d.current.NextString(); has {
			return s, true
		}
		d.current = nil
	}
}

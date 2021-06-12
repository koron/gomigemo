package inflator

type joiner struct {
	first, second Inflatable
}

// Join joins two Inflatables into one Inflatable.
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

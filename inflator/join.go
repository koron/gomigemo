package inflator

func Join(first, second Inflatable) Inflatable {
	return &joiner{first, second}
}

type joiner struct {
	first, second Inflatable
}

type joining struct {
	first      Inflator
	infratable Inflatable
	second     Inflator
}

func (j *joiner) Inflate(s string) Inflator {
	return &joining{j.first.Inflate(s), j.second, nil}
}

func (j *joining) NextString() (string, bool) {
	for {
		if j.second == nil {
			if s, has := j.first.NextString(); !has {
				return "", false
			} else {
				j.second = j.infratable.Inflate(s)
			}
		}

		if s, has := j.second.NextString(); has {
			return s, has
		}
		j.second = nil
	}
}

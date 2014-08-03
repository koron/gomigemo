package server

import "github.com/koron/gomigemo/migemo"

type Migemo struct {
	dict migemo.Dict
}

func (m *Migemo) Pattern(query string, pattern *string) error {
	p, err := migemo.Pattern(m.dict, query)
	if err != nil {
		return err
	}
	*pattern = p
	return nil
}

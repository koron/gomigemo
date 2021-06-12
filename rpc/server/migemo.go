package server

import "github.com/koron/gomigemo/migemo"

// Migemo is migemo RPC server implementation.
type Migemo struct {
	dict migemo.Dict
}

// Pattern generates a migemo pattern from query.
func (m *Migemo) Pattern(query string, pattern *string) error {
	p, err := migemo.Pattern(m.dict, query)
	if err != nil {
		return err
	}
	*pattern = p
	return nil
}

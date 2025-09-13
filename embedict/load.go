package embedict

import (
	"github.com/koron/gomigemo/migemo"
)

// Load loads embedded migemo.Dict.
func Load() (migemo.Dict, error) {
	return migemo.LoadAssets(&assets{})
}

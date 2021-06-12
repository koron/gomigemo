package embedict

import (
	"github.com/koron/gomigemo/migemo"
)

// Load loads embeded migemo.Dict.
func Load() (migemo.Dict, error) {
	return migemo.LoadAssets(&assets{})
}

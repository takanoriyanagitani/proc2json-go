package proc2json

import (
	"encoding/json"
)

type Proc struct {
	Environ   []string `json:"environ"`
	Cmdline   []string `json:"cmdline"`
	SessionId string   `json:"sessionid"`
}

func Convert(p *Proc) ([]byte, error) {
	return json.Marshal(p)
}

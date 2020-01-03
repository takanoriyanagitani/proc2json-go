package proc2json

import (
	"encoding/json"
)

// Proc represents proc info.
type Proc struct {
	Environ   []string `json:"environ"`
	Cmdline   []string `json:"cmdline"`
	SessionId string   `json:"sessionid"`
}

// Convert is a function to convert Proc object to json.
func Convert(p *Proc) ([]byte, error) {
	return json.Marshal(p)
}

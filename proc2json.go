package proc2json

// Proc represents proc info.
type Proc struct {
	Environ   []string `json:"environ"`
	Cmdline   []string `json:"cmdline"`
	SessionId string   `json:"sessionid"`
	Status    string   `json:"status"`
	Limits    string   `json:"limits"`
}

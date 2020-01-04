package proc2json

import (
	"testing"

	"encoding/json"
)

func TestConvert(t *testing.T) {
	p := Proc{
		Environ: []string{
			"LANG=en_US.UTF-8",
			"TERM=screen",
			"SHELL=/bin/bash",
		},
		Cmdline: []string{
			"vim",
			"proc2json.go",
		},
		SessionId: "5",
		Status:    "Name: vim\nnonvoluntary_ctxt_switches: 42\n",
		Limits:    "Limit SoftLimit HardLimit Units\nMax cpu time unlimited unlimited seconds\n",
	}
	j, e := Convert(&p)
	switch e {
	case nil:
		break
	default:
		t.Errorf("Marshal error: %v\n", e)
	}
	a := Proc{}
	em := json.Unmarshal(j, &a)
	switch em {
	case nil:
		break
	default:
		t.Errorf("Unmarshal error: %v\n", em)
	}

	switch len(a.Environ){
		case 3: break
		default: t.Errorf("Unexpected environ count: %v\n", len(a.Environ))
	}

	if "LANG=en_US.UTF-8" != a.Environ[0] {
		t.Errorf("LANG mismatch.\n")
	}

	if "TERM=screen" != a.Environ[1] {
		t.Errorf("TERM mismatch.\n")
	}

	if "SHELL=/bin/bash" != a.Environ[2] {
		t.Errorf("SHELL mismatch.\n")
	}

	switch len(a.Cmdline){
		case 2: break
		default: t.Errorf("cmdline item count mismatch.\n")
	}

	if "vim" != a.Cmdline[0] {
		t.Errorf("exename mismatch.\n")
	}

	if "proc2json.go" != a.Cmdline[1] {
		t.Errorf("arg1 mismatch.\n")
	}

	if "5" != a.SessionId {
		t.Errorf("sessionid mismatch.\n")
	}

	if "Name: vim\nnonvoluntary_ctxt_switches: 42\n" != a.Status {
		t.Errorf("Status mismatch.\n")
	}

	if "Limit SoftLimit HardLimit Units\nMax cpu time unlimited unlimited seconds\n" != a.Limits {
		t.Errorf("Limit mismatch.\n")
	}
}

package proc2json

import (
	"bytes"
	"fmt"
	"testing"
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
	}
	j, e := Convert(&p)
	switch e {
	case nil:
		break
	default:
		t.Errorf("Marshal error: %v\n", e)
	}
	expected := []byte(fmt.Sprintf(
		`{"environ":["%s","%s","%s"],"cmdline":["%s","%s"],"sessionid":"5"}`,
		"LANG=en_US.UTF-8",
		"TERM=screen",
		"SHELL=/bin/bash",
		"vim",
		"proc2json.go",
	))
	switch bytes.Equal(j, expected) {
	case true:
		break
	default:
		t.Errorf(
			"bytes mismatch. expected: %v, actual: %v\n",
			string(expected),
			string(j),
		)
	}
}

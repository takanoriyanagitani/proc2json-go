package proc2json

import (
	"testing"
)

func TestEnviron2strings(t *testing.T) {
	environ := []byte("TERM=screen\u0000SHELL=/bin/bash\u0000")
	a := Environ2strings(environ)
	e := []string{
		"TERM=screen",
		"SHELL=/bin/bash",
	}
	switch len(a) == len(e) {
	case true:
		break
	case false:
		t.Errorf("elements count mismatch. expected: %v, actual: %v\n", len(e), len(a))
	}
	for i := 0; i < len(a); i++ {
		switch a[i] == e[i] {
		case true:
			break
		case false:
			t.Errorf("item mismatch.\n")
		}
	}
}

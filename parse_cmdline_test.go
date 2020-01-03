package proc2json

import (
  "testing"
)

func TestCmdline2strings(t *testing.T){
  cmdline := []byte("vim\u0000proc2json.go\u0000")
  a := Cmdline2strings(cmdline)
  e := []string{
    "vim",
    "proc2json.go",
  }
  switch len(a) == len(e) {
    case true:  break
    case false: t.Errorf("elements count mismatch. expected: %v, actual: %v\n", len(e), len(a))
  }
  for i := 0; i < len(a); i++ {
    switch a[i] == e[i] {
      case true:  break
      case false: t.Errorf("item mismatch.\n")
    }
  }
}

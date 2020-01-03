package proc2json

import (
  "testing"
  "fmt"
  "bytes"
)

func TestConvert(t *testing.T){
  p := Proc{
    Environ: []string{
      "LANG=en_US.UTF-8",
      "TERM=screen",
      "SHELL=/bin/bash",
    },
  }
  j, e := Convert(&p)
  switch e {
    case nil: break
    default:  t.Errorf("Marshal error: %v\n", e)
  }
  expected := []byte(fmt.Sprintf(`{"environ":["%s","%s","%s"]}`, "LANG=en_US.UTF-8", "TERM=screen", "SHELL=/bin/bash"))
  switch bytes.Equal(j, expected){
    case true: break
    default:   t.Errorf(
      "bytes mismatch. expected: %v, actual: %v\n",
      string(expected),
      string(j),
    )
  }
}

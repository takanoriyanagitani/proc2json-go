package proc2json

import (
  "bytes"
)

func Cmdline2strings(cmdline []byte) []string {
  b := bytes.Split(cmdline, []byte("\u0000"))
  l := -1+len(b)
  s := make([]string, l)
  for i := 0; i<l; i++ { s[i] = string(b[i]) }
  return s
}

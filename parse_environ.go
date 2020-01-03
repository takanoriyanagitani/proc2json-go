package proc2json

import (
  "bytes"
)

func Environ2strings(environ []byte) []string {
  b := bytes.Split(environ, []byte("\u0000"))
  l := -1+len(b)
  s := make([]string, l)
  for i := 0; i<l; i++ { s[i] = string(b[i]) }
  return s
}

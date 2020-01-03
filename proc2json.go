package proc2json

import (
  "encoding/json"
)

type Proc struct {
  Environ []byte `json:"environ"`
}

func Convert(p *Proc) ([]byte, error){
  return json.Marshal(p)
}

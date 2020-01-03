package proc2json

import (
  "encoding/json"
)

type Proc struct {
  Environ []string `json:"environ"`
}

func Convert(p *Proc) ([]byte, error){
  return json.Marshal(p)
}

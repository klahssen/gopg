package sereal

import (
	"github.com/Sereal/Sereal/Go/sereal"
)

func Marshal(payload interface{}) ([]byte, error) {
	return sereal.Marshal(payload)
}

func Unmarshal(b []byte, dest interface{}) error {
	return sereal.Unmarshal(b, dest)
}

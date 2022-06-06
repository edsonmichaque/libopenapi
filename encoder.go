package libopenapi

import "github.com/edsonmichaque/libopenapi/types"

type Encoder interface {
	Encode(*types.Spec) ([]byte, error)
}

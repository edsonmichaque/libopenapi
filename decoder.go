package libopenapi

import "github.com/edsonmichaque/libopenapi/types"

type Decoder interface {
	Decode() (*types.Spec, error)
}

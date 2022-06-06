package umbeluzi

import "github.com/edsonmichaque/umbeluzi/types"

type Decoder interface {
	Decode() (*types.Spec, error)
}

package umbeluzi

import "github.com/edsonmichaque/umbeluzi/types"

type Encoder interface {
	Encode(*types.Spec) ([]byte, error)
}

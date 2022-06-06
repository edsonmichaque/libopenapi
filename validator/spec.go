package validator

import (
	"github.com/edsonmichaque/libopenapi/types"
)

type SpecValidator struct{}

func (sv SpecValidator) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

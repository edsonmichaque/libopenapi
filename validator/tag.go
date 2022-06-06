package validator

import (
	"github.com/edsonmichaque/libopenapi/types"
)

type TagValidator struct{}

func (tv TagValidator) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

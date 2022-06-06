package validator

import (
	"github.com/edsonmichaque/libopenapi/types"
)

type Spec struct{}

func (sv Spec) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func NewSpecValidator() Spec {
	return Spec{}
}

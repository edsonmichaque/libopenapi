package validator

import (
	"errors"

	"github.com/edsonmichaque/libopenapi/types"
)

var (
	ErrNotImplemented = errors.New("libopenapi: not implemented")
)

type Validator struct {
	spec Spec
}

func (v Validator) Validate(s *types.Spec) error {
	if err := v.spec.Validate(s); err != nil {
		return err
	}

	return nil
}

func New() Validator {
	return Validator{}
}

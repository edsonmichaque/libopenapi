package validator

import (
	"github.com/edsonmichaque/libopenapi/types"
)

type License struct{}

func (tv Tag) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func NewLicenseValidator() License {
	return License{}
}

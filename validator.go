package libopenapi

import (
	"errors"

	"github.com/edsonmichaque/libopenapi/types"
)

var (
	ErrNotImplemented = errors.New("libopenapi: not implemented")
)

type operationValidator struct{}

func (ov operationValidator) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func newOperationValidator() operationValidator {
	return operationValidator{}
}

type specValidator struct{}

func (sv specValidator) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func newSpecValidator() specValidator {
	return specValidator{}
}

type licenseValidator struct{}

func (lv licenseValidator) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func newLicenseValidator() licenseValidator {
	return licenseValidator{}
}

type tagValidator struct{}

func (tv tagValidator) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func newTagValidator() tagValidator {
	return tagValidator{}
}

package validator

import (
	"errors"

	"github.com/edsonmichaque/libopenapi/types"
)

var (
	ErrNotImplemented = errors.New("libopenapi: not implemented")
)

func Tag(s *types.Spec) error {
	return ErrNotImplemented
}

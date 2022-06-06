package validator

import (
	"github.com/edsonmichaque/libopenapi/types"
)

type Tag struct{}

func (tv License) Validate(s *types.Spec) error {
	return ErrNotImplemented
}

func NewTagValidator() Tag {
	return Tag{}
}

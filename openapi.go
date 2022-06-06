package libopenapi

import (
	"github.com/edsonmichaque/libopenapi/types"
	"github.com/edsonmichaque/libopenapi/validator"
)

func New(dec Decoder, options ...Option) (*types.Spec, error) {
	validators := []func(*types.Spec) error{
		validator.Spec,
		validator.Tag,
	}

	return newSpec(dec, validators, options...)
}

func newSpec(dec Decoder, validators []func(*types.Spec) error, options ...Option) (*types.Spec, error) {
	spec, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option.Apply(spec)
	}

	for _, v := range validators {
		if err := v(spec); err != nil {
			return nil, err
		}
	}

	return spec, nil
}

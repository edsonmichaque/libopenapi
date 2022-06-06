package libopenapi

import (
	"github.com/edsonmichaque/libopenapi/types"
	"github.com/edsonmichaque/libopenapi/validator"
)

func New(dec Decoder, options ...Option) (*types.Spec, error) {
	validators := []Validator{
		validator.SpecValidator{},
		validator.TagValidator{},
	}

	return newSpec(dec, validators, options...)
}

func newSpec(dec Decoder, vs []Validator, options ...Option) (*types.Spec, error) {
	spec, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option.Apply(spec)
	}

	for _, v := range vs {
		if err := v.Validate(spec); err != nil {
			return nil, err
		}
	}

	return spec, nil
}

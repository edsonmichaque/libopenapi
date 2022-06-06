package libopenapi

import (
	"github.com/edsonmichaque/libopenapi/types"
)

func New(dec Decoder, options ...Option) (*types.Spec, error) {
	spec, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option.Apply(spec)
	}

	return spec, nil
}

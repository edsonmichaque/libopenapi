package libopenapi

import "github.com/edsonmichaque/libopenapi/types"

type Option interface {
	Apply(*doc)
}

type Validator interface {
	Validate(*types.Spec) error
}

func WithValidator(v Validator) Option {
	return withValidator{
		v: v,
	}
}

type withValidator struct {
	v Validator
}

func (wv withValidator) Apply(d *doc) {
	d.v = append(d.v, wv.v)
}

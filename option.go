package umbeluzi

import "github.com/edsonmichaque/umbeluzi/types"

type Option interface {
	Apply(*Document)
}

type validator interface {
	Validate(*types.Spec) error
}

func WithValidator(v validator) Option {
	return withValidator{
		v: v,
	}
}

type withValidator struct {
	v validator
}

func (wv withValidator) Apply(d *Document) {
	d.v = append(d.v, wv.v)
}

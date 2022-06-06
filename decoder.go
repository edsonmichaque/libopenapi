package umbeluzi

import (
	"encoding/json"
	"io"

	"github.com/edsonmichaque/umbeluzi/types"
	"gopkg.in/yaml.v3"
)

type Decoder interface {
	Decode(r io.Reader) (*types.Spec, error)
}

type yamlDecoder struct {
	fn func(io.Reader) *yaml.Decoder
}

func (j yamlDecoder) Decode(r io.Reader) (*types.Spec, error) {
	var spec types.Spec

	if err := j.fn(r).Decode(&spec); err != nil {
		return nil, err
	}

	return &spec, nil
}

func NewYAMLDecoder() yamlDecoder {
	return yamlDecoder{
		fn: yaml.NewDecoder,
	}
}

type jsonDecoder struct {
	fn func(io.Reader) *json.Decoder
}

func (j jsonDecoder) Decode(r io.Reader) (*types.Spec, error) {
	var spec types.Spec

	if err := j.fn(r).Decode(&spec); err != nil {
		return nil, err
	}

	return &spec, nil
}

func NewJSONDecoder() jsonDecoder {
	return jsonDecoder{
		fn: json.NewDecoder,
	}
}

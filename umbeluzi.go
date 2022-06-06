package umbeluzi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/edsonmichaque/umbeluzi/types"
	"gopkg.in/yaml.v3"
)

func New(data interface{}) (*types.Spec, error) {
	var fetcher Fetcher
	switch d := data.(type) {
	case string:
		fetcher = NewFSFetcher(d)

		if strings.HasPrefix(d, "http://") || strings.HasPrefix(d, "https://") {
			fetcher = NewURLFetcher(d)
		}
	case []byte:
		fetcher = NewBytesFetcher(d)
		return nil, nil
	}

	r, err := fetcher.Fetch()
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		return nil, err
	}

	var decoder Decoder = NewYAMLDecoder()

	dat := map[string]interface{}{}
	if err := yaml.Unmarshal(buf.Bytes(), &dat); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			if err := json.Unmarshal(buf.Bytes(), &dat); err != nil {
				return nil, err
			}

			decoder = NewJSONDecoder()
		}

		return nil, err
	}

	return newDocument(fetcher, decoder)
}

func NewDocument(f Fetcher, dec Decoder, options ...Option) (*types.Spec, error) {
	validators := []Option{
		WithValidator(newSpecValidator()),
	}

	return newDocument(f, dec, append(options, validators...)...)
}

func newDocument(f Fetcher, dec Decoder, options ...Option) (*types.Spec, error) {
	r, err := f.Fetch()
	if err != nil {
		return nil, err
	}

	spec, err := dec.Decode(r)
	if err != nil {
		return nil, err
	}

	doc := Document{
		s: spec,
	}

	for _, option := range options {
		option.Apply(&doc)
	}

	return spec, nil
}

var (
	ErrEnd = errors.New("umbeluzi: end")
)

type Document struct {
	s *types.Spec
	v []validator
}

type schemas map[string]*types.Schema

func (s schemas) Find(name string) (*types.Schema, error) {
	return nil, errors.New("umbeluzi: not found")
}

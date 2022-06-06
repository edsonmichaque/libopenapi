package oas3

type MediaType struct {
	Schema   *Schema            `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  interface{}        `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]Example `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding *Encoding          `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

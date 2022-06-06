package oas3

type Info struct {
	Title         string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description   string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermOfService string   `json:"termOfService,omitempty" yaml:"termOfService,omitempty"`
	Contact       *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License       *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version       string   `json:"version,omitempty" yaml:"version,omitempty"`
}

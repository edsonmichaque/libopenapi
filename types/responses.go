package types

type Response struct {
	*Reference     `yaml:",inline"`
	*ResponseValue `yaml:",inline"`
}

type ResponseValue struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]Schema    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       string               `json:"links,omitempty" yaml:"links,omitempty"`
}

package types

type RequestBody struct {
	*Ref              `yaml:",inline"`
	*RequestBodyValue `yaml:",inline"`
}

type RequestBodyValue struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool                 `json:"required" yaml:"required"`
}

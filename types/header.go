package types

type Header struct {
	*Reference   `yaml:",inline"`
	*HeaderValue `yaml:",inline"`
}

type HeaderValue struct {
	Description     string      `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool        `json:"required,omitmepty" yaml:"required,omitmepty"`
	Deprecated      bool        `json:"deprecated,omitmepty" yaml:"deprecated,omitmepty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitmepty" yaml:"allowEmptyValue,omitmepty"`
	Stype           string      `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         bool        `json:"explode,omitmepty" yaml:"explode,omitmepty"`
	AllowReserved   bool        `json:"allowReserved,omitmepty" yaml:"allowReserved,omitmepty"`
	Schema          *Schema     `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty" yaml:"example,omitempty"`
}

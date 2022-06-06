package types

type Parameter struct {
	*Ref            `yaml:",inline"`
	*ParameterValue `yaml:",inline"`
}

type ParameterValue struct {
	Name            string      `json:"name,omitempty" yaml:"name,omitempty"`
	In              string      `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string      `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool        `json:"required" yaml:"required"`
	Deprecated      bool        `json:"deprecated" yaml:"deprecated"`
	AllowEmptyValue bool        `json:"allowEmptyValue" yaml:"allowEmptyValue"`
	Stype           string      `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         bool        `json:"explode" yaml:"explode"`
	AllowReserved   bool        `json:"allowReserved" yaml:"allowReserved"`
	Schema          *Schema     `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty" yaml:"example,omitempty"`
}

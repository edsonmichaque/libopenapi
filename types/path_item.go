package types

type PathItem struct {
	*Reference     `yaml:",inline"`
	*PathItemValue `yaml:",inline"`
}

type PathItemValue struct {
	Summary     string     `json:"summary" yaml:"summary"`
	Description string     `json:"description" yaml:"description"`
	Get         *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post        *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Put         *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Delete      *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch       *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
	Head        *Operation `json:"head,omitempty" yaml:"head,omitempty"`
	Options     *Operation `json:"options,omitempty" yaml:"options,omitempty"`
	Trace       *Operation `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []Server   `json:"servers,omitempty" yaml:"servers,omitempty"`
}

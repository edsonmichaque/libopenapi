package types

type Components struct {
	Schemas         map[string]Schema              `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]Response            `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]Example             `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]RequestBody         `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]Header              `json:"headers,omitempty" yaml:"headers,omitempty"`
	Links           map[string]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	SecuritySchemes map[string]SecurityScheme      `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Callbacks       map[string]map[string]PathItem `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

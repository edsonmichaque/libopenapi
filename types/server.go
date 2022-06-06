package types

type Server struct {
	URL         string                    `json:"url,omitempty" yaml:"url,omitempty"`
	Description string                    `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

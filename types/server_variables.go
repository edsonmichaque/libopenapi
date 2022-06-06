package types

type ServerVariable struct {
	Enum        []string `json:"enum,omitmepty" yaml:"enum,omitmepty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitmepty"`
	Description string   `json:"description,omitmepty" yaml:"description,omitmepty"`
}

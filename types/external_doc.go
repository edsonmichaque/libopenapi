package types

type ExternalDocs struct {
	URL         string `json:"url,omitmepty" yaml:"url,omitmepty"`
	Description string `json:"description,omitmepty" yaml:"description,omitmepty"`
}

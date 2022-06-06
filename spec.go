package oas3

type Option func(*Spec)

type Parser interface {
	Parse() (*Spec, error)
}

func New(parser Parser, options ...Option) (*Spec, error) {
	openapi, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option(openapi)
	}

	return openapi, nil
}

type Spec struct {
	OpenAPI      string                `json:"openapi" yaml:"openapi"`
	Info         Info                  `json:"info,omitempty" yaml:"info,omitempty"`
	Servers      []Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        map[string]PathItem   `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components   *Components           `json:"components,omitempty" yaml:"components,omitempty"`
	Security     []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

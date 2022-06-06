package types

type SecurityScheme struct {
	*Ref                 `yaml:",inline"`
	*SecuritySchemeValue `yaml:",inline"`
}

type SecuritySchemeValue struct {
	Type             string     `json:"type,omitempty" yaml:"type,omitempty"`
	Description      string     `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string     `json:"name,omitempty" yaml:"name,omitempty"`
	In               string     `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           string     `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     string     `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIDConnectURL string     `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}

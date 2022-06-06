package types

type Operation struct {
	Tags         []string            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string              `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs interface{}         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string              `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBody        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]Response `json:"responses,omitempty" yaml:"responses,omitempty"`
	Servers      []Server            `json:"servers,omitempty" yaml:"servers,omitempty"`
}

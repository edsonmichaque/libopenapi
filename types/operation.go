package types

import (
	"fmt"
	"strings"
)

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

func (o Operation) Validate(ctx ValidationContext) error {
	if o.OperationID == "" {
		return fmt.Errorf("paths.%s.%s.operationId must not be empty", ctx.Path, strings.ToLower(ctx.Method))
	}

	if o.Summary == "" {
		return fmt.Errorf("paths.%s.%s.summary must not be empty", ctx.Path, strings.ToLower(ctx.Method))
	}

	if o.Description == "" {
		return fmt.Errorf("paths.%s.%s.description must not be empty", ctx.Path, strings.ToLower(ctx.Method))
	}

	return nil
}

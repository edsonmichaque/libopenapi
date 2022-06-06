package types

import (
	"fmt"
	"net/http"
)

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

func (p PathItemValue) Operations() map[string]*Operation {
	newMap := make(map[string]*Operation)

	if p.Get != nil {
		newMap[http.MethodGet] = p.Get
	}

	if p.Post != nil {
		newMap[http.MethodPost] = p.Post
	}

	if p.Put != nil {
		newMap[http.MethodPut] = p.Put
	}

	if p.Delete != nil {
		newMap[http.MethodDelete] = p.Delete
	}

	if p.Patch != nil {
		newMap[http.MethodPatch] = p.Patch
	}

	if p.Trace != nil {
		newMap[http.MethodTrace] = p.Trace
	}

	if p.Options != nil {
		newMap[http.MethodOptions] = p.Options
	}

	return newMap
}

func (p PathItemValue) Validate(ctx ValidationContext) error {
	if p.Summary == "" {
		return fmt.Errorf("paths.%s.summary is empty", ctx.Path)
	}

	if p.Description == "" {
		return fmt.Errorf("paths.%s.description is empty", ctx.Path)
	}

	if p.Get != nil {
		newCtx := ValidationContext{
			Method: http.MethodGet,
			Path:   ctx.Path,
		}

		if err := p.Get.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Head != nil {
		newCtx := ValidationContext{
			Method: http.MethodHead,
			Path:   ctx.Path,
		}

		if err := p.Head.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Options != nil {
		newCtx := ValidationContext{
			Method: http.MethodOptions,
			Path:   ctx.Path,
		}

		if err := p.Options.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Post != nil {
		newCtx := ValidationContext{
			Method: http.MethodPost,
			Path:   ctx.Path,
		}

		if err := p.Post.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Put != nil {
		newCtx := ValidationContext{
			Method: http.MethodPut,
			Path:   ctx.Path,
		}

		if err := p.Put.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Delete != nil {
		newCtx := ValidationContext{
			Method: http.MethodDelete,
			Path:   ctx.Path,
		}

		if err := p.Delete.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Trace != nil {
		newCtx := ValidationContext{
			Method: http.MethodTrace,
			Path:   ctx.Path,
		}

		if err := p.Trace.Validate(newCtx); err != nil {
			return err
		}
	}

	if p.Patch != nil {
		newCtx := ValidationContext{
			Method: http.MethodPatch,
			Path:   ctx.Path,
		}

		if err := p.Patch.Validate(newCtx); err != nil {
			return err
		}
	}

	return nil
}

type ValidationContext struct {
	Method string
	Path   string
}

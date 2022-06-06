package types

import (
	"errors"
	"fmt"
)

type Schema struct {
	*Reference   `yaml:",inline"`
	*SchemaValue `yaml:",inline"`
}

type Reference struct {
	Ref *string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type SchemaValue struct {
	ID                   string            `json:"$id,omitempty" yaml:"$id,omitempty"`
	Schema               string            `json:"$schema,omitempty" yaml:"$schema,omitempty"`
	Type                 string            `json:"type,omitempty" yaml:"type,omitempty"`
	Format               string            `json:"format,omitempty" yaml:"format,omitempty"`
	Title                string            `json:"title,omitempty" yaml:"title,omitempty"`
	Description          string            `json:"description,omitempty" yaml:"description,omitempty"`
	Default              interface{}       `json:"default,omitempty" yaml:"default,omitempty"`
	AllOf                []*Schema         `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf                []*Schema         `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf                []*Schema         `json:"anyOf,omitempty" yaml:"any,omitempty"`
	Not                  *Schema           `json:"not,omitempty" yaml:"not,omitempty"`
	Discrimitator        *Discriminator    `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	Required             []string          `json:"required,omitempty" yaml:"required,omitempty"`
	AdditionalProperties interface{}       `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Example              interface{}       `json:"example,omitempty" yaml:"example,omitempty"`
	Enum                 []interface{}     `json:"enum,omitempty" yaml:"enum,omitempty"`
	Pattern              string            `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MultipleOf           *float64          `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum              *float64          `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum     *float64          `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum              *float64          `json:"minimum,omitempty" yaml:"minimum"`
	ExclusiveMinimum     *float64          `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength            *int64            `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength            *int64            `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxItems             *int64            `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems             *int64            `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems          *int64            `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties        *int64            `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties        *int64            `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Properties           map[string]Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	PatternProperties    map[string]Schema `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
	XML                  *XML              `json:"xml,omitempty" yaml:"xml,omitempty"`
	Deprecated           bool              `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	ReadOnly             bool              `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly            bool              `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
}

type Discriminator struct {
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

func builtinFormats() map[string][]string {
	return map[string][]string{
		"string": {
			"email",
			"date",
			"date-time",
			"password",
			"byte",
			"binary",
			"duration",

			"idn-email",
			"hostname",
			"idn-hostname",

			"ipv4",
			"ipv6",

			"uuid",

			"uri",
			"uri-reference",

			"iri",
			"iri-reference",

			"regex",
			"uri-template",

			"json-pointer",
			"relative-json-pointer",
		},
		"integer": {
			"int32",
			"int64",
		},
		"number": {
			"float",
			"double",
		},
	}
}

func builtinTypes() []string {
	return []string{
		"string",
		"number",
		"integer",
		"null",
		"object",
		"array",
		"boolean",
	}
}

func (s Schema) hasReference() bool {
	return s.Reference != nil
}

func (s Schema) hasValue() bool {
	return s.SchemaValue != nil
}

func (s Schema) validateType() error {
	if !contains(builtinTypes(), s.Type) {
		return fmt.Errorf("type %s is not valid", s.Type)
	}

	return nil
}

// func (s Schema) validateFormat() error {
// 	formats := builtinFormats()[s.Type]

// 	if !contains(formats, s.Format) {
// 		return fmt.Errorf("format %s is not valid", s.Type)
// 	}

// 	return nil
// }

func (s Schema) Validate() error {
	if s.hasReference() && s.hasValue() {
		return errors.New("value and reference cannot be both defined")
	}

	if err := s.validateType(); err != nil {
		return err
	}

	return nil
}

func contains(list []string, elem string) bool {
	for _, item := range list {
		if elem == item {
			return true
		}
	}

	return false
}

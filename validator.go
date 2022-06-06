package libopenapi

import "github.com/edsonmichaque/libopenapi/types"

type Validator interface {
	Validate(*types.Spec) error
}

package libopenapi

import "github.com/edsonmichaque/libopenapi/types"

type Option interface {
	Apply(*types.Spec)
}

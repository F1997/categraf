package parser

import (
	"github.com/F1997/categraf/types"
)

type Parser interface {
	Parse(input []byte, slist *types.SampleList) error
}

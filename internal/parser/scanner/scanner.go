package scanner

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// Scanner is the interface that describes a scanner.
type Scanner interface {
	HasNext() bool
	Next() token.Token
	Peek() token.Token
}

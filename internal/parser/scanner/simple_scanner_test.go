package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var _ token.Token = (*testToken)(nil)

type testToken struct {
	typ   token.Type
	value string
}

func (t testToken) Type() token.Type { return t.typ }
func (t testToken) Value() string    { return t.value }
func (t testToken) Offset() int      { return -1 }
func (t testToken) Line() int        { return -1 }
func (t testToken) Col() int         { return -1 }
func (t testToken) Length() int      { return len(t.value) }

func Test_simpleScanner(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []token.Token
	}{
		{
			"empty input",
			"",
			[]token.Token{},
		},
		{
			"single SELECT",
			"SELECT",
			[]token.Token{
				testToken{token.KeywordSelect, "SELECT"},
			},
		},
		{
			"multiple SELECT",
			"SELECT SELECT",
			[]token.Token{
				testToken{token.KeywordSelect, "SELECT"},
				testToken{token.KeywordSelect, "SELECT"},
			},
		},
		{
			"erronous multiple SELECT",
			"SELECTSELECT",
			[]token.Token{
				testToken{token.Error, "unexpected token 'SELECTSELECT'"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			sc := newSimpleScanner([]rune(tt.input))
			i := 0
			for sc.HasNext() {
				got := sc.Next()
				if got.Type() == token.EOF {
					break
				}
				assert.Equal(tt.want[i].Type(), got.Type(), "expected type '%s', but got type '%s'", tt.want[i].Type(), got.Type())
				assert.Equal(tt.want[i].Value(), got.Value(), "expected value '%s', but got value '%s'", tt.want[i].Value(), got.Value())
			}
		})
	}
}

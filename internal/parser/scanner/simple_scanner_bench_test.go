package scanner

import (
	"testing"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var result interface{}

func Benchmark_isKeyword(b *testing.B) {
	b.Run("SELECT", benchIsKeyword("SELECT"))
	b.Run("ALTER", benchIsKeyword("ALTER"))
	b.Run("UPDATE", benchIsKeyword("UPDATE"))
	b.Run("DELETE", benchIsKeyword("DELETE"))
	b.Run("FROM", benchIsKeyword("FROM"))
}

func benchIsKeyword(candidate string) func(*testing.B) {
	var r bool
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r = isKeyword(candidate)
		}
		result = r
	}
}

func BenchmarkNext(b *testing.B) {
	b.Run("SELECT", benchNextKeyword("SELECT"))
	b.Run("erronous", benchNextKeyword("FROMSELECTFOOBAR"))
}

func benchNextKeyword(keyword string) func(*testing.B) {
	var r token.Token
	rs := []rune(keyword)
	return func(b *testing.B) {
		sc := newSimpleScanner(rs)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r = sc.Next()
			sc.start, sc.pos, sc.line, sc.col = 0, 0, 0, 0 // reset position so that we don't have to create a new scanner
		}
		result = r
	}
}

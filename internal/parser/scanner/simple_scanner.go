package scanner

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

var (
	keywords     = []string{"ABORT", "ACTION", "ADD", "AFTER", "ALL", "ALTER", "ANALYZE", "AND", "AS", "ASC", "ATTACH", "AUTO INCREMENT", "BEFORE", "BEGIN", "BETWEEN", "BY", "CASCADE", "CASE", "CAST", "CHECK", "COLLATE", "COLUMN", "COMMIT", "CONFLICT", "CONSTRAINT", "CREATE", "CROSS", "CURRENT", "CURRENTDATE", "CURRENTTIME", "CURRENTTIMESTAMP", "DATABASE", "DEFAULT", "DEFERRABLE", "DEFERRED", "DELETE", "DESC", "DETACH", "DISTINCT", "DO", "DROP", "EACH", "ELSE", "END", "ESCAPE", "EXCEPT", "EXCLUDE", "EXCLUSIVE", "EXISTS", "EXPLAIN", "FAIL", "FILTER", "FIRST", "FOLLOWING", "FOR", "FOREIGN", "FROM", "FULL", "GLOB", "GROUP", "GROUPS", "HAVING", "IF", "IGNORE", "IMMEDIATE", "IN", "INDEX", "INDEXED", "INITIALLY", "INNER", "INSERT", "INSTEAD", "INTERSECT", "INTO", "IS", "ISNULL", "JOIN", "KEY", "LAST", "LEFT", "LIKE", "LIMIT", "MATCH", "NATURAL", "NO", "NOT", "NOTHING", "NOTNULL", "NULL", "NULLS", "OF", "OFFSET", "ON", "OR", "ORDER", "OTHERS", "OUTER", "OVER", "PARTITION", "PLAN", "PRAGMA", "PRECEDING", "PRIMARY", "QUERY", "RAISE", "RANGE", "RECURSIVE", "REFERENCES", "REGEXP", "REINDEX", "RELEASE", "RENAME", "REPLACE", "RESTRICT", "RIGHT", "ROLLBACK", "ROW", "ROWS", "SAVEPOINT", "SELECT", "SET", "TABLE", "TEMP", "TEMPORARY", "THEN", "TIES", "TO", "TRANSACTION", "TRIGGER", "UNBOUNDED", "UNION", "UNIQUE", "UPDATE", "USING", "VACUUM", "VALUES", "VIEW", "VIRTUAL", "WHEN", "WHERE", "WINDOW", "WITH", "WITHOUT"}
	keywordTypes = map[string]token.Type{
		"ABORT":            token.KeywordAbort,
		"ACTION":           token.KeywordAction,
		"ADD":              token.KeywordAdd,
		"AFTER":            token.KeywordAfter,
		"ALL":              token.KeywordAll,
		"ALTER":            token.KeywordAlter,
		"ANALYZE":          token.KeywordAnalyze,
		"AND":              token.KeywordAnd,
		"AS":               token.KeywordAs,
		"ASC":              token.KeywordAsc,
		"ATTACH":           token.KeywordAttach,
		"AUTO INCREMENT":   token.KeywordAutoincrement,
		"BEFORE":           token.KeywordBefore,
		"BEGIN":            token.KeywordBegin,
		"BETWEEN":          token.KeywordBetween,
		"BY":               token.KeywordBy,
		"CASCADE":          token.KeywordCascade,
		"CASE":             token.KeywordCase,
		"CAST":             token.KeywordCast,
		"CHECK":            token.KeywordCheck,
		"COLLATE":          token.KeywordCollate,
		"COLUMN":           token.KeywordColumn,
		"COMMIT":           token.KeywordCommit,
		"CONFLICT":         token.KeywordConflict,
		"CONSTRAINT":       token.KeywordConstraint,
		"CREATE":           token.KeywordCreate,
		"CROSS":            token.KeywordCross,
		"CURRENT":          token.KeywordCurrent,
		"CURRENTDATE":      token.KeywordCurrentDate,
		"CURRENTTIME":      token.KeywordCurrentTime,
		"CURRENTTIMESTAMP": token.KeywordCurrentTimestamp,
		"DATABASE":         token.KeywordDatabase,
		"DEFAULT":          token.KeywordDefault,
		"DEFERRABLE":       token.KeywordDeferrable,
		"DEFERRED":         token.KeywordDeferred,
		"DELETE":           token.KeywordDelete,
		"DESC":             token.KeywordDesc,
		"DETACH":           token.KeywordDetach,
		"DISTINCT":         token.KeywordDistinct,
		"DO":               token.KeywordDo,
		"DROP":             token.KeywordDrop,
		"EACH":             token.KeywordEach,
		"ELSE":             token.KeywordElse,
		"END":              token.KeywordEnd,
		"ESCAPE":           token.KeywordEscape,
		"EXCEPT":           token.KeywordExcept,
		"EXCLUDE":          token.KeywordExclude,
		"EXCLUSIVE":        token.KeywordExclusive,
		"EXISTS":           token.KeywordExists,
		"EXPLAIN":          token.KeywordExplain,
		"FAIL":             token.KeywordFail,
		"FILTER":           token.KeywordFilter,
		"FIRST":            token.KeywordFirst,
		"FOLLOWING":        token.KeywordFollowing,
		"FOR":              token.KeywordFor,
		"FOREIGN":          token.KeywordForeign,
		"FROM":             token.KeywordFrom,
		"FULL":             token.KeywordFull,
		"GLOB":             token.KeywordGlob,
		"GROUP":            token.KeywordGroup,
		"GROUPS":           token.KeywordGroups,
		"HAVING":           token.KeywordHaving,
		"IF":               token.KeywordIf,
		"IGNORE":           token.KeywordIgnore,
		"IMMEDIATE":        token.KeywordImmediate,
		"IN":               token.KeywordIn,
		"INDEX":            token.KeywordIndex,
		"INDEXED":          token.KeywordIndexed,
		"INITIALLY":        token.KeywordInitially,
		"INNER":            token.KeywordInner,
		"INSERT":           token.KeywordInsert,
		"INSTEAD":          token.KeywordInstead,
		"INTERSECT":        token.KeywordIntersect,
		"INTO":             token.KeywordInto,
		"IS":               token.KeywordIs,
		"ISNULL":           token.KeywordIsnull,
		"JOIN":             token.KeywordJoin,
		"KEY":              token.KeywordKey,
		"LAST":             token.KeywordLast,
		"LEFT":             token.KeywordLeft,
		"LIKE":             token.KeywordLike,
		"LIMIT":            token.KeywordLimit,
		"MATCH":            token.KeywordMatch,
		"NATURAL":          token.KeywordNatural,
		"NO":               token.KeywordNo,
		"NOT":              token.KeywordNot,
		"NOTHING":          token.KeywordNothing,
		"NOTNULL":          token.KeywordNotnull,
		"NULL":             token.KeywordNull,
		"NULLS":            token.KeywordNulls,
		"OF":               token.KeywordOf,
		"OFFSET":           token.KeywordOffset,
		"ON":               token.KeywordOn,
		"OR":               token.KeywordOr,
		"ORDER":            token.KeywordOrder,
		"OTHERS":           token.KeywordOthers,
		"OUTER":            token.KeywordOuter,
		"OVER":             token.KeywordOver,
		"PARTITION":        token.KeywordPartition,
		"PLAN":             token.KeywordPlan,
		"PRAGMA":           token.KeywordPragma,
		"PRECEDING":        token.KeywordPreceding,
		"PRIMARY":          token.KeywordPrimary,
		"QUERY":            token.KeywordQuery,
		"RAISE":            token.KeywordRaise,
		"RANGE":            token.KeywordRange,
		"RECURSIVE":        token.KeywordRecursive,
		"REFERENCES":       token.KeywordReferences,
		"REGEXP":           token.KeywordRegexp,
		"REINDEX":          token.KeywordReindex,
		"RELEASE":          token.KeywordRelease,
		"RENAME":           token.KeywordRename,
		"REPLACE":          token.KeywordReplace,
		"RESTRICT":         token.KeywordRestrict,
		"RIGHT":            token.KeywordRight,
		"ROLLBACK":         token.KeywordRollback,
		"ROW":              token.KeywordRow,
		"ROWS":             token.KeywordRows,
		"SAVEPOINT":        token.KeywordSavepoint,
		"SELECT":           token.KeywordSelect,
		"SET":              token.KeywordSet,
		"TABLE":            token.KeywordTable,
		"TEMP":             token.KeywordTemp,
		"TEMPORARY":        token.KeywordTemporary,
		"THEN":             token.KeywordThen,
		"TIES":             token.KeywordTies,
		"TO":               token.KeywordTo,
		"TRANSACTION":      token.KeywordTransaction,
		"TRIGGER":          token.KeywordTrigger,
		"UNBOUNDED":        token.KeywordUnbounded,
		"UNION":            token.KeywordUnion,
		"UNIQUE":           token.KeywordUnique,
		"UPDATE":           token.KeywordUpdate,
		"USING":            token.KeywordUsing,
		"VACUUM":           token.KeywordVacuum,
		"VALUES":           token.KeywordValues,
		"VIEW":             token.KeywordView,
		"VIRTUAL":          token.KeywordVirtual,
		"WHEN":             token.KeywordWhen,
		"WHERE":            token.KeywordWhere,
		"WINDOW":           token.KeywordWindow,
		"WITH":             token.KeywordWith,
		"WITHOUT":          token.KeywordWithout,
	}
)

func init() {
	sort.Strings(keywords) // sort keywords so that isKeyword can work properly
}

var _ Scanner = (*simpleScanner)(nil) // ensure that simpleScanner implements Scanner

func isWhitespace(r rune) bool { return unicode.IsSpace(r) }

func isKeyword(candidate string) bool {
	index := sort.SearchStrings(keywords, strings.ToUpper(candidate))
	return index != len(keywords) && keywords[index] == candidate
}

type simpleScanner struct {
	input      []rune
	start, pos int
	line, col  int
}

func New(input []rune) Scanner {
	return newSimpleScanner(input)
}

func newSimpleScanner(input []rune) *simpleScanner {
	return &simpleScanner{
		input: input,
		start: 0,
		pos:   0,
		line:  1,
		col:   1,
	}
}

func (s *simpleScanner) HasNext() bool {
	// this scanner always has a next token, but if it has reached it's EOF, the
	// next token will always be an EOF token
	return true
}

func (s *simpleScanner) Next() token.Token {
	// drain whitespaces at beginning of new token
	for !s.isEOF() && isWhitespace(s.input[s.pos]) {
		s.consumeRune()
	}
	s.start = s.pos

	next, ok := s.lookahead()
	if !ok {
		return s.token(token.EOF)
	}
	// ABCDEFGHIJKLMNOPQRSTUVW
	switch next {
	case ';':
		s.consumeRune()
		return s.token(token.StatementSeparator)
	case '|', '*', '/', '%', '<', '>', '&', '=', '!', '~':
		return s.scanOperator()
	default:
		if ('a' <= next && next <= 'w') ||
			('A' <= next && next <= 'W') {
			return s.scanKeyword()
		}
		return s.scanLiteral()
	}
}

func (s *simpleScanner) Peek() token.Token {
	start, pos, line, col := s.start, s.pos, s.line, s.col
	tk := s.Next()
	s.start, s.pos, s.line, s.col = start, pos, line, col
	return tk
}

func (s *simpleScanner) lookahead() (r rune, ok bool) {
	if s.isEOF() {
		return 0, false
	}

	return s.input[s.pos], true
}

func (s *simpleScanner) isEOF() bool {
	return s.pos >= len(s.input)
}

func (s *simpleScanner) consumeRune() {
	s.pos++
}

func (s *simpleScanner) candidate() string {
	return string(s.input[s.start:s.pos])
}

func (s *simpleScanner) errorToken(format string, args ...interface{}) token.Token {
	tk := token.New(s.line, s.col, s.start, s.pos-s.start, token.Error, fmt.Sprintf(format, args...))
	s.start = s.pos
	return tk
}

func (s *simpleScanner) token(t token.Type) token.Token {
	tk := token.New(s.line, s.col, s.start, s.pos-s.start, t, s.candidate())
	s.start = s.pos
	return tk
}

// scan functions

func (s *simpleScanner) scanOperator() token.Token {
	panic("implement me")
}

func (s *simpleScanner) scanKeyword() token.Token {
	for {
		if next, ok := s.lookahead(); ok && !isWhitespace(next) {
			s.consumeRune()
		} else {
			break
		}
	}

	candidate := s.candidate()
	if isKeyword(candidate) {
		return s.token(keywordTypes[candidate])
	}
	return s.errorToken("unexpected token '%s'", candidate)
}

func (s *simpleScanner) scanLiteral() token.Token {
	panic("implement me")
}

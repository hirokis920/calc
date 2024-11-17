package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ELLEGAL = "ELLEGAL"
	EOF     = "EOF"

	// 識別子、リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPARAN = "("
	RPARAN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

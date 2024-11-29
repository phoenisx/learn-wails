package calculator

type TokenType string

const (
	TokenNumber   TokenType = "NUMBER"
	TokenPlus     TokenType = "+"
	TokenMinus    TokenType = "-"
	TokenMultiply TokenType = "×"
	TokenDivide   TokenType = "/"
	TokenEnd      TokenType = "=" // Parsing a string or text file, `EOF` would've made sense for string end state.
)

// Token represents nodes in our tree, where Value represents the stored data to process
type Token struct {
	Type TokenType

	// We can use NullableFloat like done https://pkg.go.dev/database/sql#NullFloat64,
	// to preserve space when token is not a number
	Number float64
}

type Node struct {
	Token Token
	Left  *Node
	Right *Node
}

func (n *Node) IsNumber() bool {
	return n.Token.Type == TokenNumber
}

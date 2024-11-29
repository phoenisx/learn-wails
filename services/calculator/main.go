package calculator

import (
	"fmt"
	"learn-wails/utils"
	"strconv"
)

type Calculator struct {
	// Shunting Yard Algorithm, where the number token will always be a leaf node.
	// - https://en.wikipedia.org/wiki/Shunting_yard_algorithm
	// - https://www.youtube.com/watch?v=Wz85Hiwi5MY
	// - https://www.youtube.com/watch?v=1VjJe1PeExQ
	// Read about Expression Tree here: https://www.geeksforgeeks.org/expression-tree/
	//
	// In the above algorithm there are two steps
	// - Parsing Expression using Stack and Queue, where represents the Output queue
	// - The output queue is a Postorder tree traversal representation and can be evaluated to get the result of the expression.
	//
	// While our logic is inspired by this algorithm, it does not follow it entirely.
	// We will create an Expression Tree, but it will not support all the features offered
	// by the Shunting Yard Algorithm.
	ExpressionTree *Node
}

func New() *Calculator {
	return &Calculator{}
}

func NewToken(v string) *Token {
	number, err := strconv.ParseFloat(v, 64)

	if err == nil {
		return &Token{TokenNumber, number}
	}

	token := TokenType(v)
	return &Token{token, 0}
}

// Insert a token into the expression tree
func (calc *Calculator) Insert(in string) {
	newNode := &Node{Token: *NewToken(in)}
	if calc.ExpressionTree == nil {
		calc.ExpressionTree = newNode
	} else {
		calc.ExpressionTree = calc.insertNode(calc.ExpressionTree, newNode)
	}
}

// Remember to use the same float bit size in the result, which was used to create number values
// when creating `NewToken()`
func (calc *Calculator) Evaluate() float64 {
	return evaluateNode(calc.ExpressionTree).Value
}

func (calc *Calculator) Reset() {
	calc.ExpressionTree = nil
}

// Our logic to insert a node does not care about `percendence` of the operation
// If we do add support for Mathematical Expressions in our Calculator App, we will have to properly
// implement Shunting Yard algorithm, using a separate stack to help create a Expression Tree with proper order.
func (calc *Calculator) insertNode(current, newNode *Node) *Node {
	if !current.IsNumber() && newNode.IsNumber() {
		current.Right = newNode
	}

	if !newNode.IsNumber() {
		newNode.Left = current
		return newNode // New Node becomes the root
	}

	return current
}

func getIdentityValue(t TokenType) float64 {
	switch t {
	case TokenPlus, TokenMinus:
		return 0
	case TokenMultiply, TokenDivide:
		return 1
	}
	return 0
}

// Recursive helper to evaluate a node
func evaluateNode(node *Node) utils.NullFloat64 {
	result := utils.NullFloat64{Valid: false}
	if node == nil {
		return result
	}

	result.Valid = true
	// Leaf node (number)
	if node.IsNumber() {
		result.Value = node.Token.Number
		return result
	}

	// Internal node (operator)
	leftVal := evaluateNode(node.Left)
	right := evaluateNode(node.Right)

	rightVal := right.Value
	if !right.Valid {
		rightVal = getIdentityValue(node.Token.Type)
	}

	switch node.Token.Type {
	case TokenPlus:
		result.Value = leftVal.Value + rightVal
		return result
	case TokenMinus:
		result.Value = leftVal.Value - rightVal
		return result
	case TokenMultiply:
		result.Value = leftVal.Value * rightVal
		return result
	case TokenDivide:
		result.Value = leftVal.Value / rightVal
		return result
	default:
		panic(fmt.Sprintf("Unknown operator: %v", node.Token))
	}
}

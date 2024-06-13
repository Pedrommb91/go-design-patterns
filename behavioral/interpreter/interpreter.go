package interpreter

// Context holds information that's global to the interpreter
type Context struct {
	VariableValues map[string]int
}

// Expression is the interface for all expressions
type Expression interface {
	Interpret(ctx *Context) int
}

// VariableExpression represents a variable in the language
type VariableExpression struct {
	Name string
}

// Interpret returns the value of the variable from the context
func (v *VariableExpression) Interpret(ctx *Context) int {
	return ctx.VariableValues[v.Name]
}

// AddExpression represents addition between two expressions
type AddExpression struct {
	Left  Expression
	Right Expression
}

// Interpret calculates the sum of the left and right expressions' values
func (a *AddExpression) Interpret(ctx *Context) int {
	return a.Left.Interpret(ctx) + a.Right.Interpret(ctx)
}

// SubtractExpression represents subtraction between two expressions
type SubtractExpression struct {
	Left  Expression
	Right Expression
}

// Interpret calculates the difference of the left and right expressions' values
func (s *SubtractExpression) Interpret(ctx *Context) int {
	return s.Left.Interpret(ctx) - s.Right.Interpret(ctx)
}

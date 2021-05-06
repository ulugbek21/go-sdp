package strategy

// Operator ...
type Operator interface {
	Apply(int, int) int
}

// Operation ...
type Operation struct {
	Operator Operator
}

// Operate ...
func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

// Addition ...
type Addition struct{}

// Apply ...
func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

// Multiplication ...
type Multiplication struct{}

// Apply ...
func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

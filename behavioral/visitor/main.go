package main

import (
  "fmt"
  "strings"
)

// ExpressionVisitor ...
type ExpressionVisitor interface {
  VisitDoubleExpression(de *DoubleExpression)
  VisitAdditionExpression(ae *AdditionExpression)
}

// Expression ...
type Expression interface {
  Accept(ev ExpressionVisitor)
}

// DoubleExpression ...
type DoubleExpression struct {
  value float64
}

// Accept ...
func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
  ev.VisitDoubleExpression(d)
}

// AdditionExpression ...
type AdditionExpression struct {
  left, right Expression
}

// Accept ...
func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
  ev.VisitAdditionExpression(a)
}

// ExpressionPrinter ...
type ExpressionPrinter struct {
  sb strings.Builder
}

// VisitDoubleExpression ...
func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
  e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

// VisitAdditionExpression ...
func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
  e.sb.WriteString("(")
  ae.left.Accept(e)
  e.sb.WriteString("+")
  ae.right.Accept(e)
  e.sb.WriteString(")")
}

// NewExpressionPrinter ...
func NewExpressionPrinter() *ExpressionPrinter {
  return &ExpressionPrinter{strings.Builder{}}
}

// String ...
func (e *ExpressionPrinter) String() string {
  return e.sb.String()
}

func main() {
  // 1+(2+3)
  e := &AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  ep := NewExpressionPrinter()
  ep.VisitAdditionExpression(e)
  fmt.Println(ep.String())
}
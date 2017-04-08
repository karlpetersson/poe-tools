package main

type GreaterThan struct {}
type LesserThan struct {}
type EqualTo struct {}

type Operator interface {
    eval(int, int) bool
}

func (op GreaterThan) eval(a int, b int) bool {
    return a > b
}
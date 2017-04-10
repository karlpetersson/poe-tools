package main

type Operator interface {
    eval(int, int) bool
}

type GreaterThan struct {}
type GreaterThanOrEqualTo struct {}
type LessThan struct {}
type LessThanOrEqualTo struct {}
type EqualTo struct {}

func (op LessThan) eval(a int, b int) bool {
    return a < b
}

func (op LessThanOrEqualTo) eval(a int, b int) bool {
    return a <= b
}

func (op GreaterThan) eval(a int, b int) bool {
    return a > b
}

func (op GreaterThanOrEqualTo) eval(a int, b int) bool {
    return a >= b
}
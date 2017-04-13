package main

type Operator interface {
    eval(int, int) bool
}

type GreaterThan struct {}
type GreaterThanOrEqualTo struct {}
type LessThan struct {}
type LessThanOrEqualTo struct {}
type EqualTo struct {}

type intComparator func(int)bool 

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
/*
func GreaterThanOrEqualTo(a int) intComparator {
    return func(b int)bool {
        return b >= a
    }
}*/
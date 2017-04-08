package main
import (
    "github.com/ccbrown/poe-go/api"
)

type PropertyFilter interface {
    compare(item api.Item) bool
}

type iLevelFilter struct {
    Value int
    Op Operator
}

func (f iLevelFilter) compare(item api.Item) bool {
    return f.Op.eval(f.Value, item.ItemLevel)
}
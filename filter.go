package main

import (
    "github.com/ccbrown/poe-go/api"
)

type Filter struct {
    ItemName string
    Properties []PropertyFilter
}

func (f Filter) match(item api.Item) bool {
    if f.ItemName != item.Type {
        return false
    }

    for _, prop := range f.Properties {
        if !prop.compare(item) {
            return false
        }
    }

    return true
}

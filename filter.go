package main

import (
    "github.com/ccbrown/poe-go/api"
)

type Filter struct {
    //ItemName string
    Properties []PropertyFilter
}

func (f Filter) match(item api.Item) bool {
    if len(f.Properties) == 0 {
        // empty filter
        return false;
    }

    for _, prop := range f.Properties {
        if isMatch := prop.compare(item); isMatch == false {
            return false
        }
    }

    return true
}

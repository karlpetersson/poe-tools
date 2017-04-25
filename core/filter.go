package core

import (
    "github.com/ccbrown/poe-go/api"
)

type Filter struct {
    Enabled bool
    Properties []PropertyFilter
}

type Settings struct {
    Price int
    Enabled bool
}

func (f Filter) Match(item api.Item) bool {
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

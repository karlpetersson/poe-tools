package main
import (
    "github.com/ccbrown/poe-go/api"
    "regexp"
    "strconv"
    "fmt"
)

var chaosExConv = 100
var priceRegexp = regexp.MustCompile(`~b/o (?P<Amount>\d{1,4}) (?P<Type>(?:chaos|exa))`)
var altPriceRegexp = regexp.MustCompile(`~price (?P<Amount>\d{1,4}) (?P<Type>(?:chaos|exa))`)

type PropertyFilter interface {
    compare(item api.Item) bool
}

type nameFilter struct {
    Value string
}

func (f nameFilter) compare(item api.Item) bool {
    return item.Name == f.Value
}

type typeFilter struct {
    Value string
}

func (f typeFilter) compare(item api.Item) bool {
    return item.Type == f.Value
}

type iLevelFilter struct {
    Value int
    Op Operator
}

func (f iLevelFilter) compare(item api.Item) bool {
    return f.Op.eval(item.ItemLevel, f.Value)
}

type priceFilter struct {
    Value int
    Op Operator
}

func (f priceFilter) compare(item api.Item) bool {
    match := priceRegexp.FindStringSubmatch(item.Note)
    if match == nil {
        match = altPriceRegexp.FindStringSubmatch(item.Note)
        if match == nil {
            return false
        }
    }

    result := make(map[string]string)
    for i, name := range priceRegexp.SubexpNames() {
        if i != 0 { result[name] = match[i] }
    }
    
    price, err := strconv.Atoi(result["Amount"])
    if err != nil {
        fmt.Println(err)
    }

    if result["Type"] == "exa" {
        price *= chaosExConv
    }

    return f.Op.eval(price, f.Value)
}
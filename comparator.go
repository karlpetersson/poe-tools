package main
import (
    "github.com/ccbrown/poe-go/api"
    "regexp"
    "strconv"
    "fmt"
)

var priceRegexp = regexp.MustCompile(`~b/o (?P<Amount>\d{1,4}) chaos`)
var altPriceRegexp = regexp.MustCompile(`~price (?P<Amount>\d{1,4}) chaos`)

type PropertyFilter interface {
    compare(item api.Item) bool
}

type iLevelFilter struct {
    Value int
    Op Operator
}

type priceFilter struct {
    Value int
    Op Operator
}

func (f iLevelFilter) compare(item api.Item) bool {
    return f.Op.eval(item.ItemLevel, f.Value)
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

    return f.Op.eval(price, f.Value)
}
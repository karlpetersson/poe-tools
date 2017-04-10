package main
import (
    "github.com/ccbrown/poe-go/api"
    "regexp"
    "strconv"
    "fmt"
    "strings"
)

var priceRegexp = regexp.MustCompile(`~b/o (?P<Amount>\d{1,4}) (?P<Type>[a-z]+)`)
var altPriceRegexp = regexp.MustCompile(`~price (?P<Amount>\d{1,4}) (?P<Type>[a-z]+)`)
//var altPriceRegexp = regexp.MustCompile(`~price (?P<Amount>\d{1,4}) (?P<Type>[(?:chaos|exa)])`)

type PropertyFilter interface {
    compare(item api.Item) bool
}

type nameFilter struct {
    Value string
}

func (f nameFilter) compare(item api.Item) bool {
    return strings.Trim(item.Name, "<<set:MS>><<set:M>><<set:S>>") == f.Value
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

    //if price is in chaos, don't bother to look up
    if result["Type"] == "chaos" {
        return f.Op.eval(price, f.Value)
    }

    if val, ok := defaultConversionTable[result["Type"]]; ok {
        price = int(float64(price) * val)
        /*log.Printf("val: %v", val)
        log.Printf("price: %v", price)*/
    } else {
        return false
    }

    // handle troll listings
    if (price == 0) {
        return false
    }

    return f.Op.eval(price, f.Value)
}
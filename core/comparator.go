package core
import (
    "github.com/ccbrown/poe-go/api"
    "regexp"
    "strconv"
    "fmt"
    "strings"
    "poe-indexer/ninja"
)

var priceRegexp = regexp.MustCompile(`~b/o (?P<Amount>\d{1,4}) (?P<Type>[a-z]+)`)
var altPriceRegexp = regexp.MustCompile(`~price (?P<Amount>\d{1,4}) (?P<Type>[a-z]+)`)
//var altPriceRegexp = regexp.MustCompile(`~price (?P<Amount>\d{1,4}) (?P<Type>[(?:chaos|exa)])`)

type PropertyFilter interface {
    compare(item api.Item) bool
}

type hasExplicitMod struct {
    Type ModType
    Value bool
}

func (p hasExplicitMod) compare(item api.Item) bool {
    regexp := regexp.MustCompile(getPropStr(p.Type))
    
    for _, mod := range item.ExplicitMods {
        match := regexp.FindStringSubmatch(mod)
        hasMod := (match != nil)
        if p.Value == hasMod {
            return true
        } 
    }

    return false
}

type explicitModValueFilter struct {
    Type ModType
    Value int
    Op Operator
}

func (p explicitModValueFilter) compare(item api.Item) bool {
    regexp := regexp.MustCompile(getPropStr(p.Type))
    
    for _, mod := range item.ExplicitMods {
        match := regexp.FindStringSubmatch(mod)
        if match != nil {
            result := make(map[string]string)
            for i, name := range regexp.SubexpNames() {
                if i != 0 { result[name] = match[i] }
            }

            amount, err := strconv.Atoi(result["Amount"])
            if err != nil {
                fmt.Println(err)
            }

            if p.Op.eval(amount, p.Value) {
                return true
            }
        }
    }

    return false
}

type hasImplicitMod struct {
    Type ModType
    Value bool
}

func (p hasImplicitMod) compare(item api.Item) bool {
    regexp := regexp.MustCompile(getPropStr(p.Type))
    
    for _, mod := range item.ImplicitMods {
        match := regexp.FindStringSubmatch(mod)
        hasMod := (match != nil)
        if p.Value == hasMod {
            return true
        } 
    }

    return false
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

    // handle troll listings
    if price == 0 {
        return false
    }

    //if price is in chaos, don't bother to look up
    if result["Type"] == "chaos" {
        return f.Op.eval(price, f.Value)
    }

    if val, ok := ninja.Convtable[result["Type"]]; ok {
        price = int(float64(price) * val)
        /*log.Printf("val: %v", val)
        log.Printf("price: %v", price)*/
    } else {
        return false
    }

    return f.Op.eval(price, f.Value)
}

type isCorruptedFilter struct {
    Value bool
}

func (f isCorruptedFilter) compare(item api.Item) bool {
    return item.IsCorrupted == f.Value
}

type isUniqueFilter struct {
    Value bool
}

func (f isUniqueFilter) compare(item api.Item) bool {
    return (item.FrameType == api.UniqueItemFrameType) == f.Value
}

type typeHasPrefixFilter struct {
    Value string
}

func (f typeHasPrefixFilter) compare(item api.Item) bool {
    return strings.HasPrefix(item.Type, f.Value)
}

type typeHasSuffixFilter struct {
    Value string
}

func (f typeHasSuffixFilter) compare(item api.Item) bool {
    return strings.HasSuffix(item.Type, f.Value)
}

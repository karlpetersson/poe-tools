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
/*
type propIncChargesUsed struct {
    IntComparator compare
}

type Property interface {
    eval(item api.Item) bool
}

type numValuePropFilter struct {
    Prop Property
}*/

type ModType int
const (
    HasNoSockets ModType = iota
    IncreasedFlaskChargesUsed
    PlusToLevelOfSocketedSupportGems
    AddsDexterity
    LeechAppliesInstantlyFlask
)

type hasExplicitMod struct {
    Type ModType
    Value bool
}

func getPropStr(t ModType) string {
    switch t {
        case IncreasedFlaskChargesUsed: 
            return `(?P<Amount>\d{1,2})\% increased Charges used`
        case PlusToLevelOfSocketedSupportGems:
            return `\+(?P<Amount>\d{1}) to level of Socketed Support Gems`
        case AddsDexterity:
            return `\+(?P<Amount>\d{1,3}) to Dexterity`
        case HasNoSockets:
            return `Has no Sockets`
        case LeechAppliesInstantlyFlask:
            return `Leech applies instantly during Flask effect`
        default:
            return "Unknown property"
    }
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

/*func propTypeToStr(t PropType) string {
    regexp := regexp.MustCompile(getPropStr(t))
    
    match := regexp.FindStringSubmatch(item.Note)
    if match == nil {
}

func (p numValuePropFilter) compare(item api.Item) bool {
    propStr = propTypeToStr
    for prop, _ := range item.Properties {
        if prop.Name == p.name {
            if p.Comparator(p.Value, prop.Values[0])
        } 
    }
}*/

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

    if val, ok := defaultConversionTable[result["Type"]]; ok {
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

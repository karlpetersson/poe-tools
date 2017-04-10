package main

func loadFilters() []Filter {
    filters := make([]Filter, 10)
    
    opalRing := Filter{"Opal Ring", []PropertyFilter{}}
    opalRing.Properties = append(opalRing.Properties, iLevelFilter{84, GreaterThanOrEqualTo{}})
    opalRing.Properties = append(opalRing.Properties, priceFilter{80, LessThanOrEqualTo{}})    
    filters = append(filters, opalRing)

    aKey := Filter{"Ancient Reliquary Key", []PropertyFilter{}}
    aKeyFilter := priceFilter{60, LessThanOrEqualTo{}}
    aKey.Properties = append(aKey.Properties, aKeyFilter)
    filters = append(filters, aKey)

    exa := Filter{"Exalted Orb", []PropertyFilter{}}
    exaFilter := priceFilter{100, LessThanOrEqualTo{}}
    exa.Properties = append(exa.Properties, exaFilter)
    filters = append(filters, exa)

    return filters
}
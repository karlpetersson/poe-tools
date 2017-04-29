package main

import (
    "log"
    "poe-indexer/core"
)

func loadFilters() []core.Filter {
    filters := make([]core.Filter, 0)

    filters = append(filters, core.GetDivcardsFilters()...)    
    filters = append(filters, core.GetBasesFilters()...)  
    filters = append(filters, core.GetPropheciesFilters()...)  
    filters = append(filters, core.GetEssencesFilters()...)  
    filters = append(filters, core.GetGemsFilters()...)  
    filters = append(filters, core.GetCurrencyFilters()...)  
    //filters = append(filters, core.GetEarlyFilters()...)  
    filters = append(filters, core.GetVaalFilters()...)

    log.Printf("All filters loaded")

    return filters
}

package main

import (
    "log"
)

func loadFilters() []Filter {
    filters := make([]Filter, 0)

    // #1 Ancient reliquary key
    ancientReliquaryKeyFilter := Filter {
        []PropertyFilter {
            typeFilter {
                "Ancient Reliquary Key",
            },
            priceFilter {
                200,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, ancientReliquaryKeyFilter)

    // #2 Opal ring
    opalRingFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Opal Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                200,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, opalRingFilter)

    // #3 Steel ring
    steelRingFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Steel Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                50,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, steelRingFilter)

    // #4 Exalted Orb
    exaFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Exalted Orb",
            },
            priceFilter {
                80,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, exaFilter)

    //
    // DIVINATION CARDS
    
    // The Fiend
    theFiend := Filter{
        []PropertyFilter {
            typeFilter {
                "The Fiend",
            },
        },
    }
    filters = append(filters, theFiend)

    //
    // PROPHECIES
   
    // #6 Trash to treasure
    trashToTreasure := Filter{
        []PropertyFilter {
            typeFilter {
                "Trash To Treasure",
            },
            priceFilter {
                1000,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, trashToTreasure)
 
   /* opalRing := Filter{"Opal Ring", []PropertyFilter{}}
    opalRing.Properties = append(opalRing.Properties, iLevelFilter{84, GreaterThanOrEqualTo{}})
    opalRing.Properties = append(opalRing.Properties, priceFilter{80, LessThanOrEqualTo{}})    
    filters = append(filters, opalRing)

    aKey := Filter{"Ancient Reliquary Key", []PropertyFilter{}}
    aKeyFilter := priceFilter{100, LessThanOrEqualTo{}}
    aKey.Properties = append(aKey.Properties, aKeyFilter)
    filters = append(filters, aKey)*/

    /*exa := Filter{"Exalted Orb", []PropertyFilter{}}
    exaFilter := priceFilter{100, LessThanOrEqualTo{}}
    exa.Properties = append(exa.Properties, exaFilter)
    filters = append(filters, exa)*/

    log.Printf("All filters loaded")

    return filters
}
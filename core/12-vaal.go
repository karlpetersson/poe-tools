package core

func GetVaalFilters() []Filter {
    settings := map[string]Settings {
        "HandOfWisdomAndAction": { 5, true },
        "Astramentis": { 55, true },
    }

    filters := make([]Filter, 0)
    
    HandOfWisdomAndAction := Filter{
        settings["HandOfWisdomAndAction"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Hand of Wisdom and Action",
            },
            priceFilter {
                settings["HandOfWisdomAndAction"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, HandOfWisdomAndAction)

    Astramentis := Filter{
        settings["Astramentis"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Astramentis",
            },
            priceFilter {
                settings["Astramentis"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, Astramentis)

    return filters
}
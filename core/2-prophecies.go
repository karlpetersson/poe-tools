package core

func GetPropheciesFilters() []Filter {
    settings := map[string]Settings {
        "TrashToTreasure":  { 1500, true },
        "FatedConnections": { 550, true },
        "LostInPages":      { 40, true },
        "FireAndBrimstone": { 25, true },
        "JewellersTouch":   { 9, true },
        "MonstrousTreasure": { 15, true },
    }

    filters := make([]Filter, 0)
  
    MonstrousTreasure := Filter{
        settings["MonstrousTreasure"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Monstrous Treasure",
            },
            priceFilter {
                settings["MonstrousTreasure"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, MonstrousTreasure)

    TrashToTreasure := Filter{
        settings["TrashToTreasure"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Trash To Treasure",
            },
            priceFilter {
                settings["TrashToTreasure"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TrashToTreasure)
   
    LostInPages := Filter{
        settings["LostInPages"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Lost in the Pages",
            },
            priceFilter {
                settings["LostInPages"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, LostInPages)

    FireAndBrimstone := Filter{
        settings["FireAndBrimstone"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Fire and Brimstone",
            },
            priceFilter {
                settings["FireAndBrimstone"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, FireAndBrimstone)

    // Lost in the pages
    JewellersTouch := Filter{
        settings["JewellersTouch"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Jeweller's Touch",
            },
            priceFilter {
                settings["JewellersTouch"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, JewellersTouch)

    FatedConnections := Filter{
        settings["FatedConnections"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Fated Connections",
            },
            priceFilter {
                settings["FatedConnections"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, FatedConnections)

    return filters
}

       
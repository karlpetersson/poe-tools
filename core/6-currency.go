package core

func GetCurrencyFilters() []Filter {
    settings := map[string]Settings {
        "ExaltedOrb": { 50, true },
        "MirrorOfKalandra": { 4000, true },
        "Eternal Orb": { 800, true },
    }

    filters := make([]Filter, 0)

    ExaltedOrb := Filter{
        settings["ExaltedOrb"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Exalted Orb",
            },
            priceFilter {
                settings["ExaltedOrb"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, ExaltedOrb)

    MirrorOfKalandra := Filter{
        settings["MirrorOfKalandra"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Mirror of Kalandra",
            },
            priceFilter {
                settings["MirrorOfKalandra"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, MirrorOfKalandra)

    EternalOrb := Filter{
        settings["EternalOrb"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Eternal Orb",
            },
            priceFilter {
                settings["EternalOrb"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, EternalOrb)

    return filters
}
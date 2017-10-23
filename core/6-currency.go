package core

func GetCurrencyFilters() []Filter {
    settings := map[string]Settings {
        "ExaltedOrb": { 50, true },
        "MirrorOfKalandra": { 4000, true },
        "Eternal Orb": { 800, true },
        "AncientReliquaryKey": { 50, true },
        "MortalHope": { 30, true },
        "MortalIgnorance": { 15, true },
    }

    filters := make([]Filter, 0)

    AncientReliquaryKey := Filter{
        settings["AncientReliquaryKey"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Ancient Reliquary Key",
            },
            priceFilter {
                settings["AncientReliquaryKey"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, AncientReliquaryKey)

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

    MortalHope := Filter{
        settings["MortalHope"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Mortal Hope",
            },
            priceFilter {
                settings["MortalHope"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, MortalHope)

    MortalIgnorance := Filter{
        settings["MortalIgnorance"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Mortal Ignorance",
            },
            priceFilter {
                settings["MortalIgnorance"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, MortalIgnorance)

    return filters
}
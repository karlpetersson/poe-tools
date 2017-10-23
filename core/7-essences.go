package core

func GetEssencesFilters() []Filter {
    settings := map[string]Settings {
        "DeafeningEssenceOfWrath": { 5, false },
        "DeafeningEssenceOfSpite": { 8, true },
        "EssenceOfDelirium": { 15, true },
        "EssenceOfHysteria": { 10, true },
        "EssenceOfInsanity": { 10, true },
        "DeafeningEssenceOfAnger": { 5, true },
    }

    filters := make([]Filter, 0)

    DeafeningEssenceOfWrath := Filter {
        settings["DeafeningEssenceOfWrath"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Deafening Essence of Wrath",
            },
            priceFilter {
                settings["DeafeningEssenceOfWrath"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, DeafeningEssenceOfWrath)

    DeafeningEssenceOfSpite := Filter {
        settings["DeafeningEssenceOfSpite"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Deafening Essence of Spite",
            },
            priceFilter {
                settings["DeafeningEssenceOfSpite"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, DeafeningEssenceOfSpite)

    EssenceOfDelirium := Filter{
        settings["EssenceOfDelirium"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Essence of Delirium",
            },
            priceFilter {
                settings["EssenceOfDelirium"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, EssenceOfDelirium)

    EssenceOfHysteria := Filter{
        settings["EssenceOfHysteria"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Essence of Hysteria",
            },
            priceFilter {
                settings["EssenceOfHysteria"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, EssenceOfHysteria)

    EssenceOfInsanity := Filter{
        settings["EssenceOfInsanity"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Essence of Insanity",
            },
            priceFilter {
                settings["EssenceOfInsanity"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, EssenceOfInsanity)

    DeafeningEssenceOfAnger := Filter{
        settings["DeafeningEssenceOfAnger"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Deafening Essence of Anger",
            },
            priceFilter {
                settings["DeafeningEssenceOfAnger"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, DeafeningEssenceOfAnger)

    return filters
}

       
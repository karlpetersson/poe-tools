package core

func GetEssencesFilters() []Filter {
    settings := map[string]Settings {
        "DeafeningEssenceOfWrath": { 5, true },
        "DeafeningEssenceOfSpite": { 5, true },
        "EssenceOfDelirium": { 10, true },
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
                "Essence Of Delirium",
            },
            priceFilter {
                settings["EssenceOfDelirium"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, EssenceOfDelirium)


    return filters
}

       
package core

func GetOtherFilters() []Filter {
    settings := map[string]Settings {
        "SixSocketNgamahu": { 10, true },
        "i84Auxium": { 8, true },
        "PrimordialHarmony": { 15, true },
        "PrimordialMight": { 170, true },
    }

    filters := make([]Filter, 0)

    SixSocketNgamahu := Filter{
        settings["SixSocketNgamahu"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Ngamahu's Flame",
            },
            priceFilter {
                settings["SixSocketNgamahu"].Price,
                LessThanOrEqualTo{},
            },
            hasSocketsFilter {
                6,
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, SixSocketNgamahu)

    i84Auxium := Filter{
        settings["i84Auxium"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Auxium",
            },
            priceFilter {
                settings["i84Auxium"].Price,
                LessThanOrEqualTo{},
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, i84Auxium)
    
    PrimordialHarmony := Filter{
        settings["PrimordialHarmony"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Primordial Harmony",
            },
            priceFilter {
                settings["PrimordialHarmony"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, PrimordialHarmony)

    PrimordialMight := Filter{
        settings["PrimordialMight"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Primordial Might",
            },
            priceFilter {
                settings["PrimordialMight"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, PrimordialMight)

    return filters
}
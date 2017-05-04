package core

func GetBasesFilters() []Filter {
    settings := map[string]Settings {
        "i84CrystalBelt": { 200, true },
        "i84OpalRing": { 90, true },
        "i84SteelRing": { 30, true },
        
        "i84VaalRegalia": { 5, false },
        "i84HubrisCirclet": { 5, false },
        "i84SorcererGloves": { 5, false },
        "i84SorcererBoots": { 5, false },
        "i84VanguardBelt": { 5, false },

        // low tier
        "i81OpalRing": { 10, false },
        "i81SteelRing": { 10, false },
        "i81CrystalBelt": { 10, false },
        
        // super low tier
        "i78DiamondRing": { 1, false },
        "i78HarbingerBow": { 1, false },
        "i78VaalRegalia": { 1, false },
        "i78HubrisCirclet": { 1, false },
    }

    filters := make([]Filter, 0)

    i84OpalRing := Filter{
        settings["i84OpalRing"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Opal Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                settings["i84OpalRing"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
            isUniqueFilter {
                false,
            },
        },
    }
    filters = append(filters, i84OpalRing)

    i84SteelRing := Filter{
        settings["i84SteelRing"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Steel Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                settings["i84SteelRing"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
            isUniqueFilter {
                false,
            },
        },
    }
    filters = append(filters, i84SteelRing)

    i81SteelRing := Filter{
        settings["i81SteelRing"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Steel Ring",
            },
            iLevelFilter {
                81,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                settings["i81SteelRing"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
            isUniqueFilter {
                false,
            },
        },
    }
    filters = append(filters, i81SteelRing)

    i84CrystalBelt := Filter{
        settings["i84CrystalBelt"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Crystal Belt",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                settings["i84CrystalBelt"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
            isUniqueFilter {
                false,
            },
        },
    }
    filters = append(filters, i84CrystalBelt)

    i81CrystalBelt := Filter{
        settings["i81CrystalBelt"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Crystal Belt",
            },
            iLevelFilter {
                81,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                settings["i81CrystalBelt"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
            isUniqueFilter {
                false,
            },
        },
    }
    filters = append(filters, i81CrystalBelt)

    return filters
}

       
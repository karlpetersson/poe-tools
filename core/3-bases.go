package core

func GetBasesFilters() []Filter {
    settings := map[string]Settings {
        "i84CrystalBelt": { 200, true },
        "i84OpalRing": { 90, true },
        "i84SteelRing": { 30, true },
        
        "i84VaalRegalia": { 5, true },
        "i84HubrisCirclet": { 5, true },
        "i84SorcererGloves": { 5, true },
        "i84SorcererBoots": { 5, true },
        "i84VanguardBelt": { 5, true },

        // low tier
        "i81OpalRing": { 10, true },
        "i81SteelRing": { 10, true },
        "i81CrystalBelt": { 10, true },
        
        // super low tier
        "i78DiamondRing": { 1, true },
        "i78HarbingerBow": { 1, true },
        "i78VaalRegalia": { 1, true },
        "i78HubrisCirclet": { 1, true },
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

       
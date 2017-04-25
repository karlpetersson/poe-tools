package core

func GetEarlyFilters() []Filter {
    settings := map[string]Settings {
        "TheEnlightened": { 1, true },
        "Enlighten": { 5, true },
        "Empower": { 5, true },
        "TabulaRasa": { 14, true },
        "Lifesprig": { 2, true },
    }

    filters := make([]Filter, 0)

    TheEnlightened := Filter{
        settings["TheEnlightened"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Enlightened",
            },
            priceFilter {
                settings["TheEnlightened"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheEnlightened)

    Enlighten := Filter{
        settings["Enlighten"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Enlighten Support",
            },
            priceFilter {
                settings["Enlighten"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, Enlighten)

    Empower := Filter{
        settings["Empower"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Empower Support",
            },
            priceFilter {
                settings["Empower"].Price,
                LessThanOrEqualTo{},
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, Empower)

    TabulaRasa := Filter{
        settings["TabulaRasa"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Tabula Rasa",
            },
            priceFilter {
                settings["TabulaRasa"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TabulaRasa)
    
    Lifesprig := Filter{
        settings["Lifesprig"].Enabled,
        []PropertyFilter {
            nameFilter {
                "Lifesprig",
            },
            priceFilter {
                settings["Lifesprig"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, Lifesprig)

    //TODO: 6-linked items
    //TODO: skill gems?
    //TODO: low ilvl-bases
    //TODO: 5-linked items
    //TODO: 

    return filters
}

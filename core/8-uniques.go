package core

func GetUniquesFilters() []Filter {
    settings := map[string]Settings {
        "GrandSpectrum": { 15, true },
    }

    filters := make([]Filter, 0)
    
    GrandSpectrum := Filter{
        settings["GrandSpectrum"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Viridian Jewel",
            },
            priceFilter {
                settings["GrandSpectrum"].Price,
                LessThanOrEqualTo{},
            },
            nameFilter {
                "Grand Spectrum",
            },
        },
    }
    filters = append(filters, GrandSpectrum)

    return filters
}
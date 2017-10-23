package core

func GetDivcardsFilters() []Filter {
    settings := map[string]Settings {
        "TheFiend":         { 500, true },
        "TheDoctor":        { 800, true },
        "HouseOfMirrors":   { 1000, true },
        "TheQueen":         { 55, true },
        "TheImmortal":      { 80, true },
        "TheOffering":      { 25, true },
        "AbandonedWealth":  { 40, true },
        "TheDragonsHeart":  { 20, true },
        "TheSpark":         { 0, false },
        "MawrBlaidd":       { 0, false },
        "WealthAndPower":   { 0, false },
        "TheEthereal":      { 0, false },
        "ThePolymath":      { 0, false },
        "HuntersReward":    { 0, false },        
        "LastHope":         { 0, false },
        "TheArtist":        { 0, false },
    }
       
    filters := make([]Filter, 0)

    TheQueen := Filter{
        settings["TheQueen"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Queen",
            },
            priceFilter {
                settings["TheQueen"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheQueen)

    TheImmortal := Filter{
        settings["TheImmortal"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Immortal",
            },
            priceFilter {
                settings["TheImmortal"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheImmortal)

    TheOffering := Filter{
        settings["TheOffering"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Offering",
            },
            priceFilter {
                settings["TheOffering"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheOffering)

    TheFiend := Filter{
        settings["TheFiend"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Fiend",
            },
            priceFilter {
                settings["TheFiend"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheFiend)

    TheDoctor := Filter{
        settings["TheDoctor"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Doctor",
            },
            priceFilter {
                settings["TheDoctor"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheDoctor)

    HouseOfMirrors := Filter{
        settings["HouseOfMirrors"].Enabled,
        []PropertyFilter {
            typeFilter {
                "House Of Mirrors",
            },
            priceFilter {
                settings["HouseOfMirrors"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, HouseOfMirrors)

    AbandonedWealth := Filter{
        settings["AbandonedWealth"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Abandoned Wealth",
            },
            priceFilter {
                settings["AbandonedWealth"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, AbandonedWealth)

    TheDragonsHeart := Filter{
        settings["TheDragonsHeart"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Dragon's Heart",
            },
            priceFilter {
                settings["TheDragonsHeart"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheDragonsHeart)

    TheSpark := Filter{
        settings["TheSpark"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Spark and the Flame",
            },
            priceFilter {
                settings["TheSpark"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheSpark)

    MawrBlaidd := Filter{
        settings["MawrBlaidd"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Mawr Blaidd",
            },
            priceFilter {
                settings["MawrBlaidd"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, MawrBlaidd)

    WealthAndPower := Filter{
        settings["WealthAndPower"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Wealth and Power",
            },
            priceFilter {
                settings["WealthAndPower"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, WealthAndPower)

    TheEthereal := Filter{
        settings["TheEthereal"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Ethereal",
            },
            priceFilter {
                settings["TheEthereal"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheEthereal)

    LastHope := Filter{
        settings["LastHope"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Last Hope",
            },
            priceFilter {
                settings["LastHope"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, LastHope)

    HuntersReward := Filter{
        settings["HuntersReward"].Enabled,
        []PropertyFilter {
            typeFilter {
                "Hunter's Reward",
            },
            priceFilter {
                settings["HuntersReward"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, HuntersReward)

    ThePolymath := Filter{
        settings["ThePolymath"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Polymath",
            },
            priceFilter {
                settings["ThePolymath"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, ThePolymath)

    TheArtist := Filter{
        settings["TheArtist"].Enabled,
        []PropertyFilter {
            typeFilter {
                "The Artist",
            },
            priceFilter {
                settings["TheArtist"].Price,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, TheArtist)

    return filters
}


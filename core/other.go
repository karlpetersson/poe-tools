package core

func GetOtherFilters() []Filter {
    /*prices := map[string]int {
        true,
        "fiend": 500,
    }*/

    filters := make([]Filter, 0)

    test := Filter {
        true,
        []PropertyFilter {
            nameFilter {
                "Eye of Chayula",
            },
            hasImplicitMod {
                MoreFrenzyCharges,
                true,
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, test)

    test2 := Filter {
        true,
        []PropertyFilter {
            nameFilter {
                "Eye of Chayula",
            },
            hasImplicitMod {
                AdditionalCurse,
                true,
            },
            priceFilter {
                500,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, test2)

    test3 := Filter {
        true,
        []PropertyFilter {
            nameFilter {
                "Bisco's Collar",
            },
            hasImplicitMod {
                MoreFrenzyCharges,
                true,
            },
            priceFilter {
                75,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, test3)

    test4 := Filter {
        true,
        []PropertyFilter {
            nameFilter {
                "Bisco's Collar",
            },
            hasImplicitMod {
                AdditionalCurse,
                true,
            },
            priceFilter {
                500,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, test4)

    // expensive stuff
    vesselOfVinktar := Filter {
        true,
        []PropertyFilter {
            nameFilter {
                "Vessel of Vinktar",
            },
            hasExplicitMod {
                LeechAppliesInstantlyFlask,
                true,
            },
        },
    }
    filters = append(filters, vesselOfVinktar)

    // PRICEFIX THESE 
    eyeOfChayula := Filter {
        true,
        []PropertyFilter {
            nameFilter {
                "Eye of Chayula",
            },
            isCorruptedFilter {
                false,
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, eyeOfChayula)



    // #1 Ancient reliquary key
    ancientReliquaryKeyFilter := Filter {
        true,
        []PropertyFilter {
            typeFilter {
                "Ancient Reliquary Key",
            },
            priceFilter {
                50,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, ancientReliquaryKeyFilter)

   
    // #4 Exalted Orb
    exaFilter := Filter{
        true,
        []PropertyFilter {
            typeFilter {
                "Exalted Orb",
            },
            priceFilter {
                50,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, exaFilter)

    // Jewels
    // Red Nightmare
    theRedNightmare := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Red Nightmare",
            },
            priceFilter {
                300,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theRedNightmare)
    theBlueNightmare := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Blue Nightmare",
            },
            priceFilter {
                300,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theBlueNightmare)
    theGreenNightmare := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Green Nightmare",
            },
            priceFilter {
                220,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theGreenNightmare)

    theAnimaStone := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Anima Stone",
            },
            priceFilter {
                230,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theAnimaStone)
   
    primordialMight := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Primordial Might",
            },
            priceFilter {
                200,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, primordialMight)

    energyFromWithin := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Energy From Within",
            },
            priceFilter {
                130,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, energyFromWithin)

    theRedDream := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Red Dream",
            },
            priceFilter {
                90,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theRedDream)

    theBlueDream := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Blue Dream",
            },
            priceFilter {
                75,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theBlueDream)

    theGreenDream := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "The Green Dream",
            },
            priceFilter {
                25,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theGreenDream)

    grandSpectrum := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Grand Spectrum",
            },
            typeFilter {
                "Viridian Jewel",
            },
            priceFilter {
                25,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, grandSpectrum)

    // Uniques
    bereksRespite := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Berek's Respite",
            },
            priceFilter {
                50,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, bereksRespite)

    shavronnesRevelation := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Shavronne's Revelation",
            },
            priceFilter {
                140,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, shavronnesRevelation)

    vollsDevotion := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Voll's Devotion",
            },
            priceFilter {
                45,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, vollsDevotion)

    astramentis := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Astramentis",
            },
            priceFilter {
                40,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, astramentis)
   
    headhunter := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Headhunter",
            },
            priceFilter {
                2200,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, headhunter)

    kiaras := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Kiara's Determination",
            },
            priceFilter {
                80,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, kiaras)

    skyforth := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Skyforth",
            },
            priceFilter {
                2200,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, skyforth)

    shavs := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Shavronne's Wrappings",
            },
            priceFilter {
                350,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, shavs)

    tukohamas := Filter{
        true,
        []PropertyFilter {
            nameFilter {
                "Tukohama's Fortress",
            },
            priceFilter {
                350,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, tukohamas)

    // Gems
    enlighten := Filter{
        true,
        []PropertyFilter {
            typeFilter {
                "Enlighten Support",
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, enlighten)

    // Leagestones
    cartosLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Cartographer's Ambush Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                60,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, cartosLeaguestone)

    // Leagestones
    singularAmbushLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Singular Ambush Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, singularAmbushLeaguestone)

    // Leagestones
    singularPerandusLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Singular Perandus Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                90,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, singularPerandusLeaguestone)

    // Leagestones
    cartosPerandusLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Cartographer's Perandus Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                90,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, cartosPerandusLeaguestone)

    // Leagestones
    chayulaLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasSuffixFilter {
                "Breach Leaguestone of Dreaming",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                50,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, chayulaLeaguestone)

    // Leagestones
    plentifulBreachLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Plentiful Breach Leaguestone of Splinters",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                15,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, plentifulBreachLeaguestone)

    // Leagestones
    endbringingLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Endbringing Beyond Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, endbringingLeaguestone)

    // Leagestones
    terrorLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Talisman Leaguestone of Terror",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, terrorLeaguestone)

    deliriousLeaguestone := Filter{
        true,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Delirious Essence Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                40,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, deliriousLeaguestone)

    return filters
}
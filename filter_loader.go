package main

import (
    "log"
)

func loadFilters() []Filter {
    filters := make([]Filter, 0)

    test := Filter {
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

    // The Queen
    theQueen := Filter{
        []PropertyFilter {
            typeFilter {
                "The Queen",
            },
            priceFilter {
                55,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theQueen)

    // The Immortal
    theImmortal := Filter{
        []PropertyFilter {
            typeFilter {
                "The Immortal",
            },
            priceFilter {
                80,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theImmortal)

    // The Offering
    theOffering := Filter{
        []PropertyFilter {
            typeFilter {
                "The Offering",
            },
            priceFilter {
                30,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theOffering)

    // fun stuff 
    deafeningEssenceOfWrath := Filter {
        []PropertyFilter {
            typeFilter {
                "Deafening Essence of Wrath",
            },
            priceFilter {
                7,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, deafeningEssenceOfWrath)

    deafeningEssenceOfSpite := Filter {
        []PropertyFilter {
            typeFilter {
                "Deafening Essence of Spite",
            },
            priceFilter {
                10,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, deafeningEssenceOfSpite)

    shriekingSpite := Filter {
        []PropertyFilter {
            typeFilter {
                "Shrieking Essence of Spite",
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, shriekingSpite)


    theEnlightened := Filter{
        []PropertyFilter {
            typeFilter {
                "The Enlightened",
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theEnlightened)

    theValkyrie := Filter{
        []PropertyFilter {
            typeFilter {
                "The Valkyrie",
            },
            priceFilter {
                5,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theValkyrie)

    bowDream := Filter{
        []PropertyFilter {
            typeFilter {
                "Bowyer's Dream",
            },
            priceFilter {
                10,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, bowDream)

   brittleEmperor := Filter{
        []PropertyFilter {
            typeFilter {
                "The Brittle Emperor",
            },
            priceFilter {
                2,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, brittleEmperor)

    

    // #1 Ancient reliquary key
    ancientReliquaryKeyFilter := Filter {
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

    // #2 Opal ring
    opalRingFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Opal Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                85,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, opalRingFilter)

    // #3 Steel ring
    steelRingFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Steel Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                30,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, steelRingFilter)

    // Crystal Belt
    crystalBelt := Filter{
        []PropertyFilter {
            typeFilter {
                "Crystal Belt",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
            priceFilter {
                500,
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
    filters = append(filters, crystalBelt)

    // #4 Exalted Orb
    exaFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Exaltdepoe Orb",
            },
            priceFilter {
                50,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, exaFilter)

    //
    // DIVINATION CARDS
    
    // The Fiend
    theFiend := Filter{
        []PropertyFilter {
            typeFilter {
                "The Fiend",
            },
            priceFilter {
                500,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theFiend)

    // The Doctor
    theDoctor := Filter{
        []PropertyFilter {
            typeFilter {
                "The Doctor",
            },
            priceFilter {
                700,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theDoctor)

    // House of Mirrors
    houseOfMirrors := Filter{
        []PropertyFilter {
            typeFilter {
                "House Of Mirrors",
            },
            priceFilter {
                750,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, houseOfMirrors)


    // Abandoned Wealth
    abandonedWealth := Filter{
        []PropertyFilter {
            typeFilter {
                "Abandoned Wealth",
            },
            priceFilter {
                35,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, abandonedWealth)


    // The Dragon's Heart
    theDragonsHeart := Filter{
        []PropertyFilter {
            typeFilter {
                "The Dragon's Heart",
            },
            priceFilter {
                30,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theDragonsHeart)

    // The Spark And The Flame
    theSpark := Filter{
        []PropertyFilter {
            typeFilter {
                "The Spark and the Flame",
            },
            priceFilter {
                25,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theSpark)

    // Mawr Blaidd
    mawrBlaidd := Filter{
        []PropertyFilter {
            typeFilter {
                "Mawr Blaidd",
            },
            priceFilter {
                15,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, mawrBlaidd)

    // Wealth and Power
    wealthAndPower := Filter{
        []PropertyFilter {
            typeFilter {
                "Wealth and Power",
            },
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, wealthAndPower)

    // The Ethereal
    theEthereal := Filter{
        []PropertyFilter {
            typeFilter {
                "The Ethereal",
            },
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theEthereal)

    // The Ethereal
    lastHope := Filter{
        []PropertyFilter {
            typeFilter {
                "Last Hope",
            },
            priceFilter {
                10,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, lastHope)

    // The Ethereal
    huntersReward := Filter{
        []PropertyFilter {
            typeFilter {
                "Hunter's Reward",
            },
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, huntersReward)

    // The Ethereal
    thePolymath := Filter{
        []PropertyFilter {
            typeFilter {
                "The Polymath",
            },
            priceFilter {
                15,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, thePolymath)

    // The Ethereal
    theArtist := Filter{
        []PropertyFilter {
            typeFilter {
                "The Artist",
            },
            priceFilter {
                15,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theArtist)


    //
    // PROPHECIES
   
    // #6 Trash to treasure
    trashToTreasure := Filter{
        []PropertyFilter {
            typeFilter {
                "Trash To Treasure",
            },
            priceFilter {
                4000,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, trashToTreasure)
   
    // Lost in the pages
    lostInPages := Filter{
        []PropertyFilter {
            typeFilter {
                "Lost in the Pages",
            },
            priceFilter {
                40,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, lostInPages)

    fireAndBrimstone := Filter{
        []PropertyFilter {
            typeFilter {
                "Fire and Brimstone",
            },
            priceFilter {
                25,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, fireAndBrimstone)

    monstrousTreasure := Filter{
        []PropertyFilter {
            typeFilter {
                "Monstrous Treasure",
            },
            priceFilter {
                30,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, monstrousTreasure)

    // Lost in the pages
    jewellersTouch := Filter{
        []PropertyFilter {
            typeFilter {
                "The Jeweller's Touch",
            },
            priceFilter {
                9,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, jewellersTouch)

    fatedConnections := Filter{
        []PropertyFilter {
            typeFilter {
                "Fated Connections",
            },
            priceFilter {
                500,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, fatedConnections)

    //
    // ESSENCES
    // Delirium
    essenceDelirium := Filter{
        []PropertyFilter {
            typeFilter {
                "Essence Of Delirium",
            },
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, essenceDelirium)

    // Jewels
    // Red Nightmare
    theRedNightmare := Filter{
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

    log.Printf("All filters loaded")

    return filters
}
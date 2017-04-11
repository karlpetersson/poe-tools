package main

import (
    "log"
)

func loadDataFilters() []Filter {
    filters := make([]Filter, 0)

    // fun stuff 
    eyeOfChayula := Filter {
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Eye of Chayula",
            },
            isCorruptedFilter {
                false,
            },
        },
    }
    filters = append(filters, eyeOfChayula)

    theEnlightened := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Enlightened",
            },
        },
    }
    filters = append(filters, theEnlightened)

    theValkyrie := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Valkyrie",
            },
        },
    }
    filters = append(filters, theValkyrie)

    bowDream := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Bowyer's Dream",
            },
        },
    }
    filters = append(filters, bowDream)

   brittleEmperor := Filter{
    realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Brittle Emperor",
            },
        },
    }
    filters = append(filters, brittleEmperor)

    

    // #1 Ancient reliquary key
    ancientReliquaryKeyFilter := Filter {
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Ancient Reliquary Key",
            },
        },
    }
    filters = append(filters, ancientReliquaryKeyFilter)

    // #2 Opal ring
    opalRingFilter := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Opal Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, opalRingFilter)

    // #3 Steel ring
    steelRingFilter := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Steel Ring",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, steelRingFilter)

    // Crystal Belt
    crystalBelt := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Crystal Belt",
            },
            iLevelFilter {
                84,
                GreaterThanOrEqualTo{},
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
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Exalted Orb",
            },
        },
    }
    filters = append(filters, exaFilter)

    //
    // DIVINATION CARDS
    
    // The Fiend
    theFiend := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Fiend",
            },
        },
    }
    filters = append(filters, theFiend)

    // The Doctor
    theDoctor := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Doctor",
            },
        },
    }
    filters = append(filters, theDoctor)

    // House of Mirrors
    houseOfMirrors := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "House Of Mirrors",
            },
        },
    }
    filters = append(filters, houseOfMirrors)

    // The Immortal
    theImmortal := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Immortal",
            },
        },
    }
    filters = append(filters, theImmortal)

    // The Queen
    theQueen := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Queen",
            },
        },
    }
    filters = append(filters, theQueen)

    // Abandoned Wealth
    abandonedWealth := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Abandoned Wealth",
            },
        },
    }
    filters = append(filters, abandonedWealth)

    // The Offering
    theOffering := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Offering",
            },
        },
    }
    filters = append(filters, theOffering)

    // The Dragon's Heart
    theDragonsHeart := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Dragon's Heart",
            },
        },
    }
    filters = append(filters, theDragonsHeart)

    // The Spark And The Flame
    theSpark := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Spark and the Flame",
            },
        },
    }
    filters = append(filters, theSpark)

    // Mawr Blaidd
    mawrBlaidd := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Mawr Blaidd",
            },
        },
    }
    filters = append(filters, mawrBlaidd)

    // Wealth and Power
    wealthAndPower := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Wealth and Power",
            },
        },
    }
    filters = append(filters, wealthAndPower)

    // The Ethereal
    theEthereal := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Ethereal",
            },
        },
    }
    filters = append(filters, theEthereal)

    // The Ethereal
    lastHope := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Last Hope",
            },
        },
    }
    filters = append(filters, lastHope)

    // The Ethereal
    huntersReward := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Hunter's Reward",
            },
        },
    }
    filters = append(filters, huntersReward)

    // The Ethereal
    thePolymath := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Polymath",
            },
        },
    }
    filters = append(filters, thePolymath)

    // The Ethereal
    theArtist := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Artist",
            },
        },
    }
    filters = append(filters, theArtist)


    //
    // PROPHECIES
   
    // #6 Trash to treasure
    trashToTreasure := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Trash To Treasure",
            },
        },
    }
    filters = append(filters, trashToTreasure)
   
    // Lost in the pages
    lostInPages := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Lost in the Pages",
            },
        },
    }
    filters = append(filters, lostInPages)

    fireAndBrimstone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Fire and Brimstone",
            },
        },
    }
    filters = append(filters, fireAndBrimstone)

    monstrousTreasure := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Monstrous Treasure",
            },
        },
    }
    filters = append(filters, monstrousTreasure)

    // Lost in the pages
    jewellersTouch := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "The Jeweller's Touch",
            },
        },
    }
    filters = append(filters, jewellersTouch)

    fatedConnections := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Fated Connections",
            },
        },
    }
    filters = append(filters, fatedConnections)

    //
    // ESSENCES
    // Delirium
    essenceDelirium := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Essence Of Delirium",
            },
        },
    }
    filters = append(filters, essenceDelirium)

    // Jewels
    // Red Nightmare
    theRedNightmare := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Red Nightmare",
            },
        },
    }
    filters = append(filters, theRedNightmare)
    theBlueNightmare := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Blue Nightmare",
            },
        },
    }
    filters = append(filters, theBlueNightmare)
    theGreenNightmare := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Green Nightmare",
            },
        },
    }
    filters = append(filters, theGreenNightmare)

    theAnimaStone := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Anima Stone",
            },
        },
    }
    filters = append(filters, theAnimaStone)
   
    primordialMight := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Primordial Might",
            },
        },
    }
    filters = append(filters, primordialMight)

    energyFromWithin := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Energy From Within",
            },
        },
    }
    filters = append(filters, energyFromWithin)

    theRedDream := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Red Dream",
            },
        },
    }
    filters = append(filters, theRedDream)

    theBlueDream := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Blue Dream",
            },
        },
    }
    filters = append(filters, theBlueDream)

    theGreenDream := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "The Green Dream",
            },
        },
    }
    filters = append(filters, theGreenDream)

    grandSpectrum := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Grand Spectrum",
            },
            typeFilter {
                "Viridian Jewel",
            },
        },
    }
    filters = append(filters, grandSpectrum)

    // Uniques
    bereksRespite := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Berek's Respite",
            },
        },
    }
    filters = append(filters, bereksRespite)

    shavronnesRevelation := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Shavronne's Revelation",
            },
        },
    }
    filters = append(filters, shavronnesRevelation)

    vollsDevotion := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Voll's Devotion",
            },
        },
    }
    filters = append(filters, vollsDevotion)

    astramentis := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Astramentis",
            },
        },
    }
    filters = append(filters, astramentis)
   
    headhunter := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Headhunter",
            },
        },
    }
    filters = append(filters, headhunter)

    kiaras := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Kiara's Determination",
            },
        },
    }
    filters = append(filters, kiaras)

    skyforth := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Skyforth",
            },
        },
    }
    filters = append(filters, skyforth)

    shavs := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Shavronne's Wrappings",
            },
        },
    }
    filters = append(filters, shavs)

    tukohamas := Filter{
        realValue: 0,
        []PropertyFilter {
            nameFilter {
                "Tukohama's Fortress",
            },
        },
    }
    filters = append(filters, tukohamas)

    // Gems
    enlighten := Filter{
        realValue: 0,
        []PropertyFilter {
            typeFilter {
                "Enlighten Support",
            },
        },
    }
    filters = append(filters, enlighten)

    // Leagestones
    cartosLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Cartographer's Ambush Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, cartosLeaguestone)

    // Leagestones
    singularAmbushLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Singular Ambush Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, singularAmbushLeaguestone)

    // Leagestones
    singularPerandusLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Singular Perandus Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, singularPerandusLeaguestone)

    // Leagestones
    cartosPerandusLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Cartographer's Perandus Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, cartosPerandusLeaguestone)

    // Leagestones
    chayulaLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasSuffixFilter {
                "Breach Leaguestone of Dreaming",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, chayulaLeaguestone)

    // Leagestones
    plentifulBreachLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Plentiful Breach Leaguestone of Splinters",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, plentifulBreachLeaguestone)

    // Leagestones
    endbringingLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Endbringing Beyond Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, endbringingLeaguestone)

    // Leagestones
    terrorLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Talisman Leaguestone of Terror",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, terrorLeaguestone)

    deliriousLeaguestone := Filter{
        realValue: 0,
        []PropertyFilter {
            typeHasPrefixFilter {
                "Delirious Essence Leaguestone",
            },
            iLevelFilter {
                69,
                GreaterThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, deliriousLeaguestone)

    log.Printf("All filters loaded")

    return filters
}
package main

import (
    "log"
)

func loadFilters() []Filter {
    filters := make([]Filter, 0)

    // #1 Ancient reliquary key
    ancientReliquaryKeyFilter := Filter {
        []PropertyFilter {
            typeFilter {
                "Ancient Reliquary Key",
            },
            priceFilter {
                55,
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
                100,
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
                40,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, steelRingFilter)

    // #4 Exalted Orb
    exaFilter := Filter{
        []PropertyFilter {
            typeFilter {
                "Exalted Orb",
            },
            priceFilter {
                90,
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
                200,
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
                200,
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
                200,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, houseOfMirrors)

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

    // The Immortal
    theQueen := Filter{
        []PropertyFilter {
            typeFilter {
                "The Queen",
            },
            priceFilter {
                60,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theQueen)

    // Abandoned Wealth
    abandonedWealth := Filter{
        []PropertyFilter {
            typeFilter {
                "Abandoned Wealth",
            },
            priceFilter {
                45,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, abandonedWealth)

    // The Offering
    theOffering := Filter{
        []PropertyFilter {
            typeFilter {
                "The Offering",
            },
            priceFilter {
                45,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theOffering)

    // The Dragon's Heart
    theDragonsHeart := Filter{
        []PropertyFilter {
            typeFilter {
                "The Dragon's Heart",
            },
            priceFilter {
                35,
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
                30,
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
                20,
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
    theEnlightened := Filter{
        []PropertyFilter {
            typeFilter {
                "The Enlightened",
            },
            priceFilter {
                8,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, theEnlightened)


    //
    // PROPHECIES
   
    // #6 Trash to treasure
    trashToTreasure := Filter{
        []PropertyFilter {
            typeFilter {
                "Trash To Treasure",
            },
            priceFilter {
                1000,
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
                30,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, lostInPages)

    // #6 Trash to treasure
    fatedConnections := Filter{
        []PropertyFilter {
            typeFilter {
                "Fated Connections",
            },
            priceFilter {
                200,
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
                200,
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
                200,
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
                200,
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
                200,
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
                150,
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
                100,
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
                80,
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
            priceFilter {
                20,
                LessThanOrEqualTo{},
            },
        },
    }
    filters = append(filters, grandSpectrum)

    /* opalRing := Filter{"Opal Ring", []PropertyFilter{}}
    opalRing.Properties = append(opalRing.Properties, iLevelFilter{84, GreaterThanOrEqualTo{}})
    opalRing.Properties = append(opalRing.Properties, priceFilter{80, LessThanOrEqualTo{}})    
    filters = append(filters, opalRing)

    aKey := Filter{"Ancient Reliquary Key", []PropertyFilter{}}
    aKeyFilter := priceFilter{100, LessThanOrEqualTo{}}
    aKey.Properties = append(aKey.Properties, aKeyFilter)
    filters = append(filters, aKey)*/

    /*exa := Filter{"Exalted Orb", []PropertyFilter{}}
    exaFilter := priceFilter{100, LessThanOrEqualTo{}}
    exa.Properties = append(exa.Properties, exaFilter)
    filters = append(filters, exa)*/

    log.Printf("All filters loaded")

    return filters
}
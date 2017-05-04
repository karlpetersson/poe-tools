package core

func GetCorruptionFilters() []Filter {
    settings := map[string]Settings {
        "CurseAmulet": { 15, true },
        "EnduranceChargeBelt": { 5, true },
        "SixLinkedPlusGems": { 1000, true },
    }

    filters := make([]Filter, 0)

    CurseAmulet := Filter {
        settings["CurseAmulet"].Enabled,
        []PropertyFilter {
            priceFilter {
                settings["CurseAmulet"].Price,
                LessThanOrEqualTo{},
            },
            baseTypeFilter {
                BaseTypeAmulet{},
            },
            filterOr {
                []PropertyFilter {
                    hasImplicitMod {
                        AdditionalCurse,
                        true,
                    },
                },
            },
        },
    }
    filters = append(filters, CurseAmulet)

    EnduranceChargeBelt := Filter {
        settings["EnduranceChargeBelt"].Enabled,
        []PropertyFilter {
            priceFilter {
                settings["EnduranceChargeBelt"].Price,
                LessThanOrEqualTo{},
            },
            baseTypeFilter {
                BaseTypeBelt{},
            },
            filterOr {
                []PropertyFilter {
                    hasImplicitMod {
                        MoreEnduranceCharges,
                        true,
                    },
                },
            },
        },
    }
    filters = append(filters, EnduranceChargeBelt)

    SixLinkedPlusGems := Filter{
        settings["SixLinkedPlusGems"].Enabled,
        []PropertyFilter {
            baseTypeFilter {
                BaseTypeArmor{},
            },
            priceFilter {
                settings["SixLinkedPlusGems"].Price,
                LessThanOrEqualTo{},
            },
            hasImplicitMod {
                PlusToLevelOfSocketedGems,
                true,
            },
            sixLinkedFilter {},
        },
    }
    filters = append(filters, SixLinkedPlusGems)

    return filters
}
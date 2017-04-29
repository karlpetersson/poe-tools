package core

type BaseType interface {
    parseType(typeString string) bool
}

func _parseType(t string, names []string) bool {
    for _, name := range names {
        if t == name {
            return true
        }
    }
    return false
}

type BaseTypeAmulet struct {}
type BaseTypeBelt struct {}
type BaseTypeArmor struct {}

func (t BaseTypeAmulet) parseType(typeString string) bool {
    names := []string {
        "Citrine Amulet",
        "Onyx Amulet",
        "Lapis Amulet",
        "Coral Amulet",
        "Paua Amulet",
        "Amber Amulet",
        "Jade Amulet",
        "Gold Amulet",
        "Agate Amulet",
        "Turquoise Amulet",
        "Marble Amulet",
        "Blue Pearl Amulet",
    }
    return _parseType(typeString, names)
}

func (t BaseTypeBelt) parseType(typeString string) bool {
    names := []string {
        "Chain Belt",
        "Rustic Sash",
        "Heavy Belt",
        "Leather Belt",
        "Cloth Belt",
        "Studded Belt",
        "Vanguard Belt",
        "Crystal Belt",
    }
    return _parseType(typeString, names)
}

func (t BaseTypeArmor) parseType(typeString string) bool {
    names := []string {
        "Simple Robe",
        "Vaal Regalia",
        "Astral Plate",
    }
    return _parseType(typeString, names)
}
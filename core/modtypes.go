package core

type ModType int
const (
    HasNoSockets ModType = iota
    IncreasedFlaskChargesUsed
    PlusToLevelOfSocketedSupportGems
    AddsDexterity
    LeechAppliesInstantlyFlask
    MoreFrenzyCharges
    AdditionalCurse
)

func getPropStr(t ModType) string {
    switch t {
        case IncreasedFlaskChargesUsed: 
            return `(?P<Amount>\d{1,2})\% increased Charges used`
        case PlusToLevelOfSocketedSupportGems:
            return `\+(?P<Amount>\d{1}) to level of Socketed Support Gems`
        case AddsDexterity:
            return `\+(?P<Amount>\d{1,3}) to Dexterity`
        case HasNoSockets:
            return `Has no Sockets`
        case LeechAppliesInstantlyFlask:
            return `Leech applies instantly during Flask effect`
        case MoreFrenzyCharges:
            return `\+1 to Maximum Frenzy Charges`
        case AdditionalCurse:
            return `Enemies can have 1 additional Curse`
        default:
            return "Unknown property"
    }
}

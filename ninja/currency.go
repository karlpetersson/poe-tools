package ninja

import(
    "net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"

    "poe-indexer/util"
)

type CurrencyItem struct {
    CurrencyTypeName string `json:"currencyTypeName"`
    ChaosEquivalent float64 `json:"chaosEquivalent"`
}

var nameConversionTable = map[string]string{
    "Chaos Orb": "chaos",
    "Exalted Orb": "exa",
    "Divine Orb": "divine",
    "Gemcutter's Prism": "gcp",
    "Orb of Fusing": "fuse",
    "Orb of Alchemy": "alch",
    "Vaal Orb": "vaal",
    "Orb of Scouring": "scour",
    "Orb of Regret": "regret",
    "Regal Orb": "regal",
    "Cartographer's Chisel": "chisel",
}

var defaultConversionTable = map[string]float64{
    "chaos": 1,
    "exa": 93,
    "divine": 14,
    "master-sextant": 5.0,
    "journeyman-sextant": 3.0,
    "apprentice-sextant": 1.0,
    "gcp": 1.0,
    "regret": 1.4,
    "regal": 0.88,
    "scour": 0.8,
    "vaal": 0.65,
    "chisel": 0.41,
    "fuse": 0.35,
    "alch": 0.28,
    "blessed": 0.28,
    "chance": 0.22,
    "jew": 0.09,
    "chrom": 0.09,
   // "alt": 0.08,
    "aug": 0.03,
    "trans": 0.02,
    "coin": 0.01,
    "portal": 0.01,
    "wisdom": 0.004,
    "splinter-of-chayula": 1.8,
}

var Convtable = defaultConversionTable

func UseDefaultCurrencyRatios() {
    Convtable = defaultConversionTable
}

func FetchCurrencyRatios() {
    ratios := make(map[string]float64)

    resp, err := http.Get("http://api.poe.ninja/api/Data/GetCurrencyOverview?league=Legacy&date=" + util.Today())
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var itemWrapper struct {
        Items []CurrencyItem `json:"lines"`
    }

    if err := json.Unmarshal(body, &itemWrapper); err != nil {
        panic(err)
    }

    fmt.Printf("Fetching currency ratios\n")
    for _, val := range itemWrapper.Items {
        if name, ok := nameConversionTable[val.CurrencyTypeName]; ok {
            ratios[name] = val.ChaosEquivalent
        }
    }

    ratios["chaos"] = 1

    for k, v := range ratios {
        fmt.Printf("%v: %v \n", k, v)
    }

    Convtable = ratios
}
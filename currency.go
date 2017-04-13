package main

var defaultConversionTable = map[string]float64{
    "chaos": 1,
    "exa": 84,
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

func getDefaultCurrencyRatios() map[string]float64 {
    return defaultConversionTable
}
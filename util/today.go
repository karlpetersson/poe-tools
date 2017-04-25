package util

import (
    "time"
    "fmt"
)

func Today() string {
    t := time.Now()
    dateNow := t.Format("2006-01-02")
    fmt.Printf("Current date: %v\n", dateNow)
    return dateNow
}
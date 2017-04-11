package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/ccbrown/poe-go/api"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

var logFile *File
var writer *Writer

func getRecentCurrencyRatios() (string, error) {
	return "", nil
}

// To begin receiving newly updated items immediately, we need to get a recent change id. poe.ninja
// can provide us with that.
func getRecentChangeId() (string, error) {
	resp, err := http.Get("http://api.poe.ninja/api/Data/GetStats")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var stats struct {
		// There are a few more fields in the response, but we only care about this.
		NextChangeId string `json:"nextChangeId"`
	}
	if err := json.Unmarshal(body, &stats); err != nil {
		return "", err
	}

	return stats.NextChangeId, nil
}

func handleMatch(item api.Item, stash *api.Stash) {
	msg := newMessage(item, stash)

	//log.Printf("Match: name = %v, ign = %v, note = %v", item.Type, stash.LastCharacterName, item.Note)
	color.Green("-----")
	color.White(msg)
	
	if runtime.GOOS == "windows" {
		playNotification()
	}

	if err := clipboard.WriteAll(msg); err != nil {
		panic(err)
	}
}

func logMatch(item api.Item, stash *api.Stash, int price, int realPrice) {
	err := writer.Write("%v,%v,%v", item.Name, item.Type, price, realPrice)
}

// This is where we look through stashes for items of interest. For this example, we'll just log
// reliquary key activity. You might want to parse buyouts, play sounds, compose messages that you
// can send in whispers, etc.
func processStash(stash *api.Stash, filters []Filter) {
	for _, item := range stash.Items {
		if item.League != "Legacy" {
			continue
		}

		for _, filter := range filters {
			if filter.match(item) {
				handleMatch(item, stash)
			}
		}
	}
}

func main() {
	//messages := make(chan string)
	//c := startSocketServer()
	//go socketPush(c)
	if runtime.GOOS == "windows" {
		initSound()
	}
	
	log.Printf("requesting a recent change id from poe.ninja...")
	recentChangeId, err := getRecentChangeId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting with change id %v", recentChangeId)
	log.Printf("requesting current currency ratios from poe.ninja...")

	// open file for logging
    file, err = os.Create("dataLol.csv")
    checkError("Cannot create file", err)
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()

	/*currencyConvertionTable, err := getRecentCurrencyRatios()
	if err != nil {
		log.Printf("Could not get current ratios, using default")
		currencyConvertionTable = getDefaultCurrencyRatios()
	}*/
	//currencyConvertionTable := getDefaultCurrencyRatios()

	subscription := api.OpenPublicStashTabSubscription(recentChangeId)

	// If we get an interrupt signal (e.g. from control+c on the terminal), handle it gracefully.
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt)
		<-ch
		log.Printf("shutting down")
		subscription.Close()
	}()

	filters := loadDataFilters()

	log.Printf("num filters: %v", len(filters))

	// Loop forever over results.
	for result := range subscription.Channel {
		if result.Error != nil {
			log.Printf("error: %v", result.Error.Error())
			continue
		}
		for _, stash := range result.PublicStashTabs.Stashes {
			processStash(&stash, filters)
		}
	}
}

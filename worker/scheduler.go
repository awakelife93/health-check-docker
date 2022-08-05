package worker

import (
	"fmt"
	"time"
)

// * interval timer channel
var quitChannel = make(chan bool)

// * ticker
var ticker *time.Ticker

func quitIntervalAction() bool {
	fmt.Println("Call QuitIntervalAction")
	quit := <-quitChannel
	return quit
}

func clearTicker() {
	fmt.Println("Quit interval Action => ", quitIntervalAction())
	ticker = nil
}

func clearIntervalChannel() {
	quitChannel = nil
}

func StartScheduler(hour time.Duration, action func()) {
	ticker = time.NewTicker(hour * time.Hour)

	// * first call
	action()

	for {
		select {
		// case <-time.After(delay):
		case <-ticker.C:
			action()
		case <-quitChannel:
			ticker.Stop()
		}
	}
}

func ClearScheduler() {
	clearTicker()
	clearIntervalChannel()
}

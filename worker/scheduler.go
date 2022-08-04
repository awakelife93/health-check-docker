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
	fmt.Println("Quit interval Action = ", quitIntervalAction())
	ticker = nil
}

func clearIntervalChannel() {
	quitChannel = nil
}

func StartScheduler(second time.Duration, delay time.Duration, action func()) chan bool {
	var defaultSecond time.Duration = 30
	if second < defaultSecond {
		second = defaultSecond
	}

	ticker = time.NewTicker(second * time.Second)

	action()

	for {
		select {
		// case <-time.After(delay):
		case <-ticker.C:
			action()
		case <-quitChannel:
			ticker.Stop()
			return quitChannel
		}
	}
}

func ClearScheduler() {
	clearTicker()
	clearIntervalChannel()
}

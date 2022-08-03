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

func StartScheduler(second time.Duration, delay time.Duration, action func()) chan bool {
	var defaultSecond time.Duration = 30
	if second < defaultSecond {
		second = defaultSecond
	}

	ticker := time.NewTicker(second * time.Second)

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

func ClearTicker() {
	fmt.Println("Quit interval Action = ", quitIntervalAction())
	ticker = nil
}

func ClearIntervalChannel() {
	quitChannel = nil
}

func ClearScheduler() {
	ClearTicker()
	ClearIntervalChannel()
}

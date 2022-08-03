package main

import (
	"fmt"

	"github.com/awakelife93/health-check-docker/utils"
	"github.com/awakelife93/health-check-docker/worker"
)

func clear() {
	fmt.Println("Clear scheduler")
	worker.ClearScheduler()
}

func work() {
	fmt.Println("Start docker container health check.")
	exitedContainersString, error := worker.GetExitedContainers()

	if error != nil {
		fmt.Println(error.Error())
		clear()
	}

	worker.CheckExitedContainer(
		utils.GenerateExitedContainerList(exitedContainersString),
	)

	fmt.Println("End docker container health check.")
}

func start() {
	worker.StartScheduler(30, 10, work)
}

func main() {
	start()
}

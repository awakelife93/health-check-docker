package main

import (
	"fmt"
	"time"

	"github.com/awakelife93/health-check-docker/utils"
	"github.com/awakelife93/health-check-docker/worker"
)

func clear() {
	fmt.Println("Clear scheduler")
	worker.ClearScheduler()
}

func work() {
	fmt.Println("Start docker container health check.", time.Now().UTC())

	exitedContainersString, getExitedContainersError := utils.GetExitedContainers()

	if getExitedContainersError != nil {
		fmt.Println("getExitedContainersError", getExitedContainersError.Error())
		clear()
	}

	exitedContainers, generateExitedContainerListError := utils.GenerateExitedContainerList(exitedContainersString)

	if generateExitedContainerListError != nil {
		fmt.Println("generateExitedContainerListError", generateExitedContainerListError.Error())
		clear()
	}

	exitedContainerReportError := worker.ExitedContainerReportAndRestart(exitedContainers)

	if exitedContainerReportError != nil {
		fmt.Println("exitedContainerReportError", exitedContainerReportError.Error())
		clear()
	}

	fmt.Println("End docker container health check.")
}

func start() {
	worker.StartScheduler(3, work)
}

func main() {
	start()
}

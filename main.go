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

	exitedContainersString, GetExitedContainersError := worker.GetExitedContainers()

	if GetExitedContainersError != nil {
		fmt.Println(GetExitedContainersError.Error())
		clear()
	}

	exitedContainers, GenerateExitedContainerListError := utils.GenerateExitedContainerList(exitedContainersString)

	if GenerateExitedContainerListError != nil {
		fmt.Println(GenerateExitedContainerListError.Error())
		clear()
	}

	output, ExitedContainerReportError := worker.ExitedContainerReport(exitedContainers)

	if ExitedContainerReportError != nil {
		fmt.Println(ExitedContainerReportError.Error())
		clear()
	}

	fmt.Println("End docker container health check.", output)
}

func start() {
	worker.StartScheduler(30, 10, work)
}

func main() {
	start()
}

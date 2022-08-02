package main

import (
	"fmt"

	"github.com/awakelife93/health-check-docker/utils"
	"github.com/awakelife93/health-check-docker/worker"
)

func start() {
	fmt.Println("Start docker container health check.")

	exitedContainersString, error := worker.GetExitedContainers()

	if error != nil {
		fmt.Println(error.Error())
	}

	worker.CheckExitedContainer(
		utils.GenerateExitedContainerList(exitedContainersString),
	)

	fmt.Println("End docker container health check.")
}

func main() {
	start()
}

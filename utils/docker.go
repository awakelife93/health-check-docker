package utils

import "strings"

func GenerateExitedContainerList(exitedContainers string) []string {
	// * (index == 0) => CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES line
	containers := strings.Split(exitedContainers, "\n")
	return containers[1 : len(containers)-1]
}

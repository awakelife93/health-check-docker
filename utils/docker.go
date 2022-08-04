package utils

import (
	"errors"
	"strings"
)

func GenerateExitedContainerList(exitedContainers string) ([]string, error) {
	// * (index == 0) => CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES line
	containers := strings.Split(exitedContainers, "\n")

	if len(containers) < 2 {
		return containers, errors.New("Exited Container Row Length < 2")
	}

	return containers[1 : len(containers)-1], nil
}

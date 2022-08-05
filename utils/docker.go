package utils

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/thoas/go-funk"
)

func GenerateExitedContainerList(exitedContainers string) ([]string, error) {
	// * (index == 0) => CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES line
	containers := strings.Split(exitedContainers, "\n")

	if len(containers) < 2 {
		return containers, errors.New("Exited Container Row Length < 2")
	}

	return containers[1 : len(containers)-1], nil
}

func GetRestartContainerIds() []string {
	// * Write your container ids here.
	return []string{}
}

func StartContainer(containerId string) (string, error) {
	const arg0 string = "docker"
	const arg1 string = "restart"

	command := exec.Command(arg0, arg1, containerId)
	var output, error bytes.Buffer
	command.Stdout = &output
	command.Stderr = &error

	commandError := command.Run()

	if commandError != nil {
		return "", commandError
	}

	if funk.IsEmpty(error.String()) {
		return "", errors.New(error.String())
	}

	return output.String(), nil
}

func GetExitedContainers() (string, error) {
	const arg0 string = "docker"
	const arg1 string = "ps"
	const arg2 string = "-a"
	const arg3 string = "-f"
	const arg4 string = "status=exited"

	command := exec.Command(arg0, arg1, arg2, arg3, arg4)
	var output, error bytes.Buffer
	command.Stdout = &output
	command.Stderr = &error

	commandError := command.Run()

	if commandError != nil {
		return "", commandError
	}

	if funk.IsEmpty(error.String()) {
		return "", errors.New(error.String())
	}

	return output.String(), nil
}

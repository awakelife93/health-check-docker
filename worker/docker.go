package worker

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func GetExitedContainers() (string, error) {
	arg0 := "docker"
	arg1 := "ps"
	arg2 := "-a"
	arg3 := "-f"
	arg4 := "status=exited"

	command := exec.Command(arg0, arg1, arg2, arg3, arg4)
	var output, error bytes.Buffer
	command.Stdout = &output
	command.Stderr = &error

	commandError := command.Run()

	if commandError != nil {
		return "", commandError
	}

	if error.String() != "" {
		return "", errors.New(error.String())
	}

	return output.String(), nil
}

func ExitedContainerReport(exitedContainers []string) (string, error) {
	for i := 0; i < len(exitedContainers); i++ {
		row := exitedContainers[i]
		response, error := SendMessage(row)

		if error != nil {
			fmt.Println("ExitedContainerReport error =>", error.Error())

			if error.Error() == "not_authed" {
				return "Failed.", error
			}
		}

		fmt.Println("Report response = ", response)
	}

	return "Completed successfully.", nil
}

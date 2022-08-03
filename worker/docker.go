package worker

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/awakelife93/health-check-docker/utils"
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

	if utils.IsStringAndNotEmpty(error.String()) {
		return "", errors.New(error.String())
	}

	return output.String(), nil
}

// todo: 체크 후 어떤 처리를 해줘야할 지? 고민해보기.
func CheckExitedContainer(exitedContainers []string) {
	for i := 0; i < len(exitedContainers); i++ {
		fmt.Println(exitedContainers[i])
	}
}

package worker

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awakelife93/health-check-docker/utils"
	"github.com/thoas/go-funk"
)

func restartContainer(row string) (bool, error) {
	requireContainerIds := utils.GetRestartContainerIds()

	if funk.IsEmpty(row) || funk.IsEmpty(requireContainerIds) {
		return false, errors.New("Failed Restart Container funk.IsEmpty(row) || funk.IsEmpty(requireContainerIds)")
	}

	column := strings.Fields(row)
	containerId := column[0]

	if funk.ContainsString(requireContainerIds, containerId) {
		output, startContainerError := utils.StartContainer(containerId)

		if startContainerError != nil {
			return false, startContainerError
		}

		fmt.Println("Success Restart Container => ", output, containerId)
		return true, nil
	}

	return false, nil
}

func ExitedContainerReportAndRestart(exitedContainers []string) error {
	for i := 0; i < len(exitedContainers); i++ {
		row := exitedContainers[i]

		slackResponse, slackApiError := SendMessage(row)

		isRestart, restartError := restartContainer(row)

		if restartError != nil {
			fmt.Println("ExitedContainerReportAndRestart Restart error =>", restartError.Error())
		}

		fmt.Println("Restart action => ", isRestart)

		if slackApiError != nil {

			if slackApiError.Error() == "not_authed" {
				return slackApiError
			}

			fmt.Println("ExitedContainerReportAndRestart Slack error =>", slackApiError.Error())
		}

		fmt.Println("Report response => ", slackResponse)
	}

	return nil
}

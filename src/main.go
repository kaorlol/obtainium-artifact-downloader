package main

import (
	"fmt"
	"time"
	"artifact-downloader/src/utils/actions"
	"artifact-downloader/src/utils/info"
)

func main() {
	println("Getting workflow latest run...")
	prevTime := time.Now()
	latestRun, err := actions.GetWorkflowLatestRun()
	if err != nil {
		fmt.Println(err)
		return
	}

	actions.DownloadArtifacts(latestRun)

	workflowInfo := info.GetInfo()
	info.UpdateInfo(info.Info{
		Status: workflowInfo.Status,
		ElapsedTime: int64(time.Since(prevTime).Seconds()),
		Workflow: info.Workflow{
			ID: workflowInfo.Workflow.ID,
			Title: workflowInfo.Workflow.Title,
		},
	})
}
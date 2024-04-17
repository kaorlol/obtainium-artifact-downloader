package main

import (
	"fmt"
	"time"

	"artifact-downloader/src/utils/modules"
	"artifact-downloader/src/utils/data"
)

func main() {
	println("Getting workflow latest run...")
	prevTime := time.Now()
	latestRun, err := modules.GetWorkflowLatestRun()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = modules.DownloadArtifacts(latestRun)
	if err != nil {
		fmt.Println(err)
		return
	}

	data.UpdateInfo(data.Info{ElapsedTime: int64(time.Since(prevTime).Seconds())})
}

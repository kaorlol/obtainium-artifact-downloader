package main

import (
	"os"
	"strings"
	"time"

	"artifact-downloader/src/data"
	"artifact-downloader/src/modules"
)

func main() {
	token := getTokenArgs()
	if token == "" {
		println("token not provided")
		return
	}

	println("Getting workflow latest run...")
	prevTime := time.Now()
	modules.SetClient(token)
	latestRun, err := modules.GetWorkflowLatestRun()
	if err != nil {
		println(err)
		return
	}

	err = modules.DownloadArtifacts(latestRun)
	if err != nil {
		println(err)
		return
	}

	data.UpdateInfo(data.Info{ElapsedTime: int64(time.Since(prevTime).Seconds())})
}

func getTokenArgs() string {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--token=") {
			return strings.TrimPrefix(args[i], "--token=")
		}
	}
	return ""
}
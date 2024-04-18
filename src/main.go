package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"artifact-downloader/src/utils/data"
	"artifact-downloader/src/utils/modules"
)

func main() {
	token := getTokenArgs()
	if token == "" {
		panic("token not provided")
	}

	println("Getting workflow latest run...")
	prevTime := time.Now()
	modules.SetClient(token)
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

func getTokenArgs() string {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--token=") {
			return strings.TrimPrefix(args[i], "--token=")
		}
	}
	return ""
}
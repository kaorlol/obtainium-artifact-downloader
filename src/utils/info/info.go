package info

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const path = "data/info.json"

type Workflow struct {
	ID    int64 `json:"id"`
	Title string `json:"title"`
}

type Info struct {
	Status    string `json:"status"`
	ElapsedTime int64 `json:"elapsed-time"`
	Workflow  Workflow `json:"workflow"`
}

func GetInfo() Info {
	file := filepath.Join(path)
	data, _ := os.ReadFile(file)

	var info Info
	json.Unmarshal(data, &info)

	return info
}

func UpdateInfo(info Info) error {
	file := filepath.Join(path)
	data, _ := json.Marshal(info)

	err := os.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetGitHubToken() string {
	tokenPat := os.Getenv("TOKEN_PAT")
	if tokenPat == "" {
		token_pat_file := filepath.Join("data", "github-token.txt")
		data, _ := os.ReadFile(token_pat_file)
		return string(data)
	}
	return tokenPat
}
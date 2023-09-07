package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	repoUrl := "https://github.com/bytedeveloperr/HNGx-Stage-1-Task"
	fileUrl := fmt.Sprintf("%s/blob/main/main.go", repoUrl)
	slackUsername := "bytdeveloper"
	track := "backend"

	resp := fmt.Sprintf(`{
		"slack_name": "%s",
		"current_day": "%s",
		"utc_time": "%s",
		"track": "%s",
		"github_file_url": "%s",
		"github_repo_url": "%s",
		"status_code": 200
		}`, slackUsername, time.Now().Weekday(), time.Now().UTC(), track, fileUrl, repoUrl,
	)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, resp)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleRequest)
	err := http.ListenAndServe(":3001", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed \n")
	} else {
		fmt.Printf("An error occured while starting the server - %s\n", err)
		os.Exit(1)
	}
}

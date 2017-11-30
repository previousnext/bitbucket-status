package main

import (
	"github.com/parnurzeal/gorequest"
	"fmt"
	"os"
)

func main() {
	status := Status{
		State: StateSuccessful,
		Key: "test123",
		Name: "Build environment",
		Url: "https://test.com",
		Description: "My cool build",
	}
	request := gorequest.New().SetBasicAuth("pnx_nick", "NkZmtJcDV6Ha5rJbgews")
	_, body, errs := request.Post("https://api.bitbucket.org/2.0/repositories/transportfornsw/tfnsw-corp/commit/88b252aa80e56fef7cdbc0a776e711c84e5d87dd/statuses/build").
		Send(status).
		End()
	fmt.Println(body)
	for _, err := range errs  {
		fmt.Println(err)
	}
	if len(errs) > 0 {
		os.Exit(1)
	}
}

// "pnx_nick", "8dkAAPt9vBbLUb6UCYJG"
// Commit hash 88b252aa80e56fef7cdbc0a776e711c84e5d87dd
// Base URL: https://bitbucket.org/transportfornsw/
// https://<bitbucket-base-url>/rest/build-status/1.0/commits/<commit-hash>
//{
//	"state": "<INPROGRESS|SUCCESSFUL|FAILED>",
//	"key": "<build-key>",
//	"name": "<build-name>",
//	"url": "<build-url>",
//	"description": "<build-description>"
//}

type State string

const (
	StateInProgress State = "INPROGRESS"
	StateSuccessful State = "SUCCESSFUL"
	StateFailed 	State = "FAILED"
	StateStopped	State = "STOPPED"
)

type Status struct {
	State       State  `json:"state"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

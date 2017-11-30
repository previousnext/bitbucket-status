package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/parnurzeal/gorequest"
)

var (
	cliStatus      = kingpin.Flag("status", "Status").Required().String()
	cliKey         = kingpin.Flag("key", "Key").Required().String()
	cliName        = kingpin.Flag("name", "Name").Required().String()
	cliUrl         = kingpin.Flag("url", "Url").Required().String()
	cliDescription = kingpin.Flag("description", "Description").Required().String()
	cliUsername    = kingpin.Flag("username", "Username").Required().Envar("BITBUCKET_STATUS_USERNAME").String()
	cliPassword    = kingpin.Flag("password", "Password").Required().Envar("BITBUCKET_STATUS_PASSWORD").String()
	cliCommitHash  = kingpin.Flag("hash", "Commit hash").Required().Envar("BITBUCKET_COMMIT").String()
	cliRepoOwner   = kingpin.Flag("owner", "Repo owner").Required().Envar("BITBUCKET_REPO_OWNER").String()
	cliRepoSlug    = kingpin.Flag("slug", "Repo slug").Required().Envar("BITBUCKET_REPO_SLUG").String()
)

const baseUrl = "https://api.bitbucket.org/2.0/repositories"

func main() {
	kingpin.Parse()

	status := Status{
		State:       *cliStatus,
		Key:         *cliKey,
		Name:        *cliName,
		Url:         *cliUrl,
		Description: *cliDescription,
	}

	request := gorequest.New().SetBasicAuth(*cliUsername, *cliPassword)
	url := fmt.Sprintf("%s/%s/%s/commit/%s/statuses/build", baseUrl, *cliRepoOwner, *cliRepoSlug, *cliCommitHash)
	_, body, errs := request.Post(url).
		Send(status).
		End()

	fmt.Println(body)
	for _, err := range errs {
		fmt.Println(err)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
}

type Status struct {
	State       string `json:"state"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

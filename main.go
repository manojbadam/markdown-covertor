package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

var logger *log.Logger

func main() {
	logger = log.StandardLogger()
	var logfile string

	// Read the default filepath from environment variable
	if os.Getenv("FILE_LOCATION") != "" {
		logfile = os.Getenv("FILE_LOCATION")
	} else {
		logfile = "sample.md"
	}

	output := ConvertGithubMarkup(logfile)
	fmt.Println(output)
}

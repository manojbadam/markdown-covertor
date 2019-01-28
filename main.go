package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	inputfile  string
	outputfile string
	logger     *log.Logger
	rootCmd    = &cobra.Command{Use: ""}
	wikiCmd    = &cobra.Command{
		Use:   "wiki",
		Short: "File to convert into wiki markdown",
		Long:  `Convert github markdown file into wiki markdown`,
		Run: func(cmd *cobra.Command, args []string) {
			output, err := ConvertGithubMarkup(inputfile)
			if err != nil {
				log.Error("Error occurred while converting the markdown", err)
				os.Exit(1)
			}
			err = ProcessOutput(output, outputfile)
			if err != nil {
				log.Error("Error occurred while writing the converted file", err)
				os.Exit(1)
			}
		},
	}
)

func main() {
	logger = log.StandardLogger()
	// var logfile string

	rootCmd.Execute()
}

// ProcessOutput writes the result to output (stdout, file etc..)
func ProcessOutput(result string, file string) error {
	if strings.ToLower(file) == "stdout" {
		fmt.Println(result)
	} else {
		err := ioutil.WriteFile(file, []byte(result), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadFile reads the github markdown file and returns slice of strings
func ReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ss := strings.Split(string(data), "\n")
	return ss, nil
}

func init() {
	wikiCmd.Flags().StringVarP(&inputfile, "input-file", "i", "sample.md", "github markdown file to convert")
	wikiCmd.Flags().StringVarP(&outputfile, "output-file", "o", "stdout", "output file of converted markdown")

	rootCmd.AddCommand(wikiCmd)
}

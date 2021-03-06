package main

import (
	"regexp"
	"strconv"
	"strings"
)

// ConvertGithubMarkup converts the github markup file to Wiki MarkDown file
func ConvertGithubMarkup(filename string) (string, error) {
	result := []string{}

	output, err := ReadFile(filename)
	if err != nil {
		return "", err
	} else {
		logger.Info("Total Lines", len(output))
		for _, v := range output {
			transformations := []func(string) string{replaceHeaders, replaceInfoMessages, replaceHyperLinks, replaceCodeBlocks, replaceBoldText, replaceInlineCodeBlocks}
			for _, fn := range transformations {
				v = fn(v)
			}
			result = append(result, v)
		}
		return strings.Join(result, "\n"), nil
	}
}

// replaceHeaders replaces the header hashes ## with h2 markup
func replaceHeaders(input string) string {
	r, _ := regexp.Compile("^[ ]{0,3}[#]{1,}")
	headerSize := r.FindString(input)
	if len(headerSize) > 0 {
		// input = len(strings.TrimSpace(headerSize))
		wikiMarkup := "h" + strconv.Itoa(len(strings.TrimSpace(headerSize))) + "."
		return strings.Replace(input, headerSize, wikiMarkup, 1)
	}
	return input
}

// replaceInfoMessages replaces the '>' markup with {info} block
func replaceInfoMessages(input string) string {
	r, _ := regexp.Compile("^[ ]{0,3}[>]{1}")
	if r.MatchString(input) {
		input = strings.TrimSpace(strings.Replace(input, ">", "", 1))
		return "{info}" + input + "{info}"
	}
	return input
}

// replaceHyperLinks replaces the hyperlinks [name](link) to [name|link]
// Regex Used to match \[(.*?)\]\((.*?)\)
func replaceHyperLinks(input string) string {
	r, _ := regexp.Compile("\\[(.*?)\\]\\((.*?)\\)")
	matches := r.FindAllString(input, -1)
	if len(matches) > 0 {
		for _, v := range matches {
			linkName := strings.Split(v[1:], "]")[0]
			linkURL := strings.Split(strings.Split(v, "(")[1], ")")[0]
			wikiMarkup := "[" + linkName + "|" + linkURL + "]"
			input = strings.Replace(input, v, wikiMarkup, -1)
		}
	}
	return input
}

// replaceCodeBlocks replaces the back ticks ``` with underscores {code} to get codeblock markup
func replaceCodeBlocks(input string) string {
	return strings.Replace(input, "```", "{code}", -1)
}

// replaceBoldText
func replaceBoldText(input string) string {
	return strings.Replace(input, "**", "*", -1)
}

// replaceInlineCodeBlocks replaces the back ticks ` with underscores _ to get italicized
func replaceInlineCodeBlocks(input string) string {
	return strings.Replace(input, "`", "_", -1)
}

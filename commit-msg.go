package main

import (
	"fmt"
	"os"
	"strings"
)

// Array of predetermined commit types.
var commitType = [4]string{
	"feature:",
	"fix:",
	"refactor:",
	"wip:",
}

// Maximum length of a given commit message.
const msgLength int = 50

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error open: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	commitMsgBytes, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid commit message. Commit messages must be in the format '<type>: <description>' and have a maximum length of 50 characters.\n")
		os.Exit(1)
	}
	commitMessage := string(commitMsgBytes)
	commitMessage = strings.Trim(commitMessage, " ")
	if len(commitMessage) > msgLength {
		fmt.Fprintf(os.Stderr, "Invalid commit message format or length. Commit messages must be in the format '<type>: <description>' and have a maximum length of 50 characters.\n")
		os.Exit(1)
	}

	firstWord := getFirstWord(commitMessage)
	var validType bool
	for i := range commitType {
		if commitType[i] == firstWord {
			validType = true
		}
	}

	if !validType {
		fmt.Fprintf(os.Stderr, "Invalid commit message format or length. Commit messages must be in the format '<type>: <description>' and have a maximum length of 50 characters.\n")
		os.Exit(1)
	}
}

// getFirstWord returns a substring ended by the first whitespace of a given string.
func getFirstWord(commitMsg string) string {
	for i := range commitMsg {
		if commitMsg[i] == ' ' {
			return commitMsg[0:i]
		}
	}
	return ""
}

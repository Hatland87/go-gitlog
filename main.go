package main

import (
	"encoding/json"
	"fmt"
	"local/logparse"
	"local/terminal"
	"net/url"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		panic("You need to input a git repo")
	}

	gitrepoURL, err := url.ParseRequestURI(os.Args[1])
	if err != nil {
		panic("Not a valid url")
	}

	fmt.Println("Cloning repo")
	_ = terminal.Run(fmt.Sprintf("git clone %s tempdir", gitrepoURL))

	fmt.Println("Collecting logs")
	rawOutput := terminal.Run("git -C tempdir log --all --numstat --no-merges --pretty=format:'%cs'")

	fmt.Println("Cleanup repo")
	_ = terminal.Run("rm -rf tempdir")

	fmt.Println("Parsing git log")
	commits := logparse.Parse(rawOutput)
	fmt.Println("Success!")

	fmt.Println("Creating json file for git log")
	file, err := os.Create("gitlog.json")
	if err != nil {
		panic(err)
	}

	formatedJSON, err := json.Marshal(commits)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(formatedJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success!")
}

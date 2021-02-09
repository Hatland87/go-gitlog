package logparse

import (
	"strconv"
	"strings"
)

type Commit struct {
	CommitDate     string
	NumFileChanged int
	NumInsertions  int
	NumDeletions   int
}

func Parse(rawOutput string) []Commit {
	commitArray := strings.Split(rawOutput, "\n\n")

	var commits []Commit

	for i := range commitArray {
		commit := getCommitsValues(commitArray[i])
		commits = append(commits, commit)
	}

	return commits
}

func getCommitsValues(commit string) Commit {
	commitArr := strings.Split(commit, "\n")

	commitDate := commitArr[0]
	numFile := len(commit) - 1

	var numInsertions int
	var numDeletions int

	for _, l := range commitArr[1:] {
		line := strings.Split(l, "\t")

		if len(line) > 1 {
			numInsertions = stringToInt(line[0])
			numDeletions = stringToInt(line[1])
		}
	}

	return Commit{
		CommitDate:     commitDate,
		NumFileChanged: numFile,
		NumInsertions:  numInsertions,
		NumDeletions:   numDeletions,
	}
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func checkIfErr(err error) {
	if err != nil {
		panic("error occurred")
	}
}

func main() {
	dir, _ := os.Getwd()
	files, err := os.ReadDir(dir)
	checkIfErr(err)

	isGitRepo := false
	for _, file := range files {
		if strings.Compare(file.Name(), ".git") == 0 {
			isGitRepo = true
			break
		}
	}

	if !isGitRepo {
		fmt.Printf("not a git repository, terminate...")
		return
	}

	cmd := exec.Command("git", "branch")
	stdout, err := cmd.Output()
	checkIfErr(err)

	gitBranches := string(stdout)
	var choseBranch string
	var currentBranch string
	branches := make([]string, 0)

	for _, branch := range strings.Split(gitBranches, "\n") {
		if len(branch) > 0 {
			branches = append(branches, branch)
			if strings.HasPrefix(branch, "*") {
				currentBranch = branch
			}
		}
	}

	fmt.Printf("current branch: %v\n", currentBranch)
	fmt.Println("choose branch to switch")
	prompt := &survey.Select{
		Message: "choose branch to switch: ",
		Options: branches,
	}
	survey.AskOne(prompt, &choseBranch)

	choseBranch = strings.TrimSpace(choseBranch)
	fmt.Printf("chose branch: %v\n", choseBranch)

	exec.Command("git", "checkout", choseBranch).Run()
	fmt.Printf("done switching to branch: %v\n", choseBranch)
}

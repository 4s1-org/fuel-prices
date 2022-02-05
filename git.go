package main

import (
	"fmt"
	"os/exec"
)

func Pull(folder string) {
	fmt.Println("Update " + folder)
	cmd := exec.Command("git", "-C", folder, "pull")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stdout))
}

func Clone(folder string, repoPath string) {
	fmt.Println("Clone " + repoPath + " to " + folder)
	cmd := exec.Command("git", "clone", "--depth", "1", "https://tankerkoenig@dev.azure.com/tankerkoenig/tankerkoenig-data/_git/tankerkoenig-data", folder)
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stdout))
}

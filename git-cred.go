package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var namespace string
	if os.Args[0] == "git" {
		if 2 >= 0 && 2 < len(os.Args) {
			namespace = os.Args[2]
		} else {
			namespace = ""
		}
	} else {
		if 1 >= 0 && 1 < len(os.Args) {
			namespace = os.Args[1]
		} else {
			namespace = ""
		}
	}
	args := []string{"config", "credential.namespace"}
	if namespace != "" {
		args = append(args, namespace)
	}
	cmd := exec.Command("git", args...)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

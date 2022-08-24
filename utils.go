package main

import (
	"bufio"
	"os"
	"os/exec"
)

func Exec(name string, args ...string) (stdout string, stderr error) {
	cmd := exec.Command(name, args...)
	bytes, stderr := cmd.Output()
	if stderr != nil {
		return
	}
	stdout = string(bytes)
	return
}

func Scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin <- scanner.Text()
	}
}

package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"runtime"
)

func main() {
	userDir, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current User Home Directory:", userDir.HomeDir)

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "dir")
	} else {
		cmd = exec.Command("sh", "-c", "ls -la")
	}
	cmd.Dir = userDir.HomeDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}

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
	os := runtime.GOOS
	fmt.Println("Detected os:", os)

	fmt.Println("Current User Home Directory:", userDir.HomeDir)

	var cmd *exec.Cmd
	var gitConfDir string

	if os == "windows" {
		gitConfDir = userDir.HomeDir
		cmd = exec.Command("cmd", "/C", "dir")
	} else {
		gitConfDir = userDir.HomeDir
		cmd = exec.Command("sh", "-c", "ls -la")
	}
	cmd.Dir = gitConfDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}

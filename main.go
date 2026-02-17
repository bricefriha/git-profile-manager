package main

import (
	"fmt"
	"os/exec"
	"os/user"
)

func main() {
	userDir, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current User Home Directory:", userDir.HomeDir)
	cmd := exec.Command("cmd", "ls", "-la")
	cmd.Dir = userDir.HomeDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}

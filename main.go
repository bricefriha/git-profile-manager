package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func main() {
	userDir, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	osDis := runtime.GOOS
	fmt.Println("Detected os:", osDis)

	fmt.Println("Current User Home Directory:", userDir.HomeDir)

	args := os.Args[1:]
	gitConfDir := userDir.HomeDir

	if args[0] == "switchto" {
		fileSource := fmt.Sprintf(".gitconfig-%s", args[1])
		sourcePath := filepath.Join(gitConfDir, fileSource)
		targetFilePath := filepath.Join(gitConfDir, ".gitconfig")

		if !fileExists(targetFilePath) {
			fmt.Fprintln(os.Stderr, "No git config available in this environment")
		}
		if !fileExists(fileSource) {
			fmt.Fprintf(os.Stderr, "%s not found", sourcePath)
			fmt.Println()
		}
	}
	// var cmd *exec.Cmd

	// if osDis == "windows" {
	// 	cmd = exec.Command("cmd", "/C", "dir")
	// } else {
	// 	gitConfDir = userDir.HomeDir
	// 	cmd = exec.Command("sh", "-c", "ls -la")
	// }
	// cmd.Dir = gitConfDir
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(string(output))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

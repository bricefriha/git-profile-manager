package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	userDir, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	args := os.Args[1:]
	gitConfDir := userDir.HomeDir

	if args[0] == "switchto" {
		profileName := args[1]
		fileSource := fmt.Sprintf(".gitconfig-%s", profileName)
		sourcePath := filepath.Join(gitConfDir, fileSource)
		targetFilePath := filepath.Join(gitConfDir, ".gitconfig")

		if !fileExists(targetFilePath) {
			fmt.Fprintln(os.Stderr, "No git config available in this environment")
			return
		}
		if !fileExists(sourcePath) {
			fmt.Fprintf(os.Stderr, "%s not found", sourcePath)
			fmt.Println()
			return
		}

		in, err := os.Open(sourcePath)
		if err != nil {
			panic(err)
		}

		out, err := os.Create(targetFilePath)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(os.Stderr, "🔄️ Switched to", profileName)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

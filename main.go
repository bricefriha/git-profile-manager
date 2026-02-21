package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	args := os.Args[1:]
	switch args[0] {

	case "-st":
		switchTo(args[1])
	case "--switchto":
		switchTo(args[1])

	case "--help":
		showHelper()
	case "-h":
		showHelper()

	default:
		fmt.Fprintln(os.Stderr, "‼️", "command '", args[0], "' not found")

	}

}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func switchTo(profileName string) {
	userDir, err := user.Current()
	gitConfDir := userDir.HomeDir

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
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

func showHelper() {
	fmt.Println("⌨️ Syntax: gpm [command] [argument]")
	fmt.Println("switchto, help")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("-st, --switchto [profileName]			switch to another profile")
	fmt.Println("-h, --help 							show gpm user guidance ")

}

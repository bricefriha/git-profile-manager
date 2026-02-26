package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {

	args := os.Args[1:]

	runCommand(args[0], args[1:])

}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func switchTo(profileName string) {
	gitConfDir := getGitConfDir()
	if strings.TrimSpace(gitConfDir) == "" {
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

func getGitConfDir() string {
	userDir, err := user.Current()
	gitConfDir := userDir.HomeDir

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return gitConfDir
}

func createProfile(profileName string) {
	userDir, err := user.Current()
	gitConfDir := userDir.HomeDir

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	destFilePath := filepath.Join(gitConfDir, fmt.Sprintf(".gitconfig-%s", profileName))
	if fileExists(destFilePath) {
		fmt.Fprintln(os.Stderr, "❌", profileName, "git profile already exist")
		return
	}

	defaultFilePath := filepath.Join(gitConfDir, ".gitconfig")

	in, err := os.ReadFile(defaultFilePath)
	if err != nil {
		panic(err)
	}

	errWf := os.WriteFile(destFilePath, in, 0644)
	if errWf != nil {
		panic(errWf)
	}
	fmt.Fprintln(os.Stderr, "✅ Just created", profileName, "git profile")
}

func showHelper() {
	fmt.Println("⌨️ Syntax: gpm [command] [argument]")
	fmt.Println("switchto, help")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("-c, --create [profileName]			    create a new profile")
	fmt.Println("-st, --switchto [profileName]			switch to another profile")
	fmt.Println("-h, --help 							show gpm user guidance ")

}

func listProfiles() {
	gitConfDir := getGitConfDir()
	if strings.TrimSpace(gitConfDir) == "" {
		return
	}

	entries, err := os.ReadDir(gitConfDir)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(entries); i++ {
		entry := entries[i]
		if strings.HasPrefix(entry.Name(), ".gitconfig-") {
			fmt.Println(strings.Split(entry.Name(), "-")[1])
		}
	}

}

func runCommand(name string, args []string) {
	switch name {

	case "-st":
		switchTo(args[0])
	case "--switchto":
		switchTo(args[0])

	case "--help":
		showHelper()
	case "-h":
		showHelper()

	case "--create":
		createProfile(args[0])
	case "-c":
		createProfile(args[0])

	case "--list":
		listProfiles()
	case "-ls":
		listProfiles()

	default:
		fmt.Fprintln(os.Stderr, "‼️", "command '", name, "' not found")

	}
}

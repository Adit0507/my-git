package main

import (
	"fmt"
	"os"

	"github.com/Adit0507/my-git.git/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "init":
		commands.Init(args)
	case "add":
		commands.Add(args)
	case "commit":
		commands.Commit(args)
	case "status":
		commands.Status(args)
	case "log":
		commands.Log(args)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: vcs <command> [args]")
	fmt.Println("\nCommands:")
	fmt.Println("  init              Initialize a new repository")
	fmt.Println("  add <file>        Add file to staging area")
	fmt.Println("  commit -m <msg>   Create a new commit")
	fmt.Println("  status            Show working tree status")
	fmt.Println("  log               Show commit history")
}

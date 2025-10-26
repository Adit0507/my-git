package commands

import (
	"fmt"
	"os"

	"github.com/Adit0507/my-git.git/internal/objects"
	"github.com/Adit0507/my-git.git/internal/repository"
	"github.com/Adit0507/my-git.git/internal/storage"
)

func Add(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: vcs add <file>")
		return
	}

	store := storage.NewStorage(".")
	index := repository.NewIndex(".")

	if err := index.Load(); err != nil {
		fmt.Printf("Error loading index: %v\n", err)
		return
	}

	for _, path := range args {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			continue
		}

		blob := objects.NewBlob(content)
		hash, err := store.WriteObject(blob)
		if err != nil {
			fmt.Printf("Error storing blob: %v\n", err)
			continue
		}

		index.Add(path, hash, "100644")
		fmt.Printf("Added %s (hash: %s)\n", path, hash[:8])
	}

	if err := index.Save(); err != nil {
		fmt.Printf("Error saving index: %v\n", err)
		return
	}
}

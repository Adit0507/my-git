package commands

import (
	"fmt"
	"os"

	"github.com/Adit0507/my-git.git/internal/objects"
	"github.com/Adit0507/my-git.git/internal/repository"
	"github.com/Adit0507/my-git.git/internal/storage"
)

func Commit(args []string) {
	if len(args) < 2 || args[0] != "-m" {
		fmt.Println("Usage: vcs commit -m <message>")
		return
	}

	message := args[1]
	repo := repository.NewRepository(".")
	store := storage.NewStorage(".")
	index := repository.NewIndex(".")

	if err := index.Load(); err != nil {
		fmt.Printf("Error loading index: %v\n", err)
		return
	}

	if len(index.Entries) == 0 {
		fmt.Println("Nothing to commit")
		return
	}

	// creatin tree from index
	tree := objects.NewTree()
	for _, entry := range index.Entries {
		tree.AddEntry(entry.Mode, entry.Path, entry.Hash)
	}

	treeHash, err := store.WriteObject(tree)
	if err != nil {
		fmt.Printf("Error storing tree: %v\n", err)
		return
	}

	parentHash, _ := repo.GetHEAD()
	author := os.Getenv("VCS_AUTHOR")
	if author == "" {
		author = "Unknown <unknown@example.com>"
	}

	// creatin commit
	commit := objects.NewCommit(treeHash, parentHash, author, message)
	commitHash, err := store.WriteObject(commit)
	if err != nil {
		fmt.Printf("error storing commit %v\n", err)
		return
	}

	// updating head
	if err := repo.UpdateHEAD(commitHash); err != nil {
		fmt.Printf("Error updating HEAD: %v\n", err)
		return
	}

	fmt.Printf("Created commit %s\n", commitHash[:8])
}

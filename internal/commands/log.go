package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/Adit0507/my-git.git/internal/objects"
	"github.com/Adit0507/my-git.git/internal/repository"
	"github.com/Adit0507/my-git.git/internal/storage"
)

func Log(args []string) {
	repo := repository.NewRepository(".")
	store := storage.NewStorage(".")

	commitHash, err := repo.GetHEAD()
	if err != nil  {
		fmt.Printf("Error reading HEAD: %v\n", err)
		return
	}

	if  commitHash == "" {
		fmt.Println("No commits yet")
		return
	}

	for commitHash != "" {
		content, objType, err := store.ReadObject(commitHash)
		if err != nil {
			fmt.Printf("Error reading commit: %v\n", err)
			return
		}

		if objType != objects.CommitType {
			fmt.Println("Invalid commit object")
			return
		}

		commit := parseCommit(content)
		fmt.Printf("commit %s\n", commitHash)
		fmt.Printf("Author: %s\n", commit.Author)
		fmt.Printf("Date:   %s\n\n", commit.Timestamp.Format(time.RFC1123))
		fmt.Printf("    %s\n\n", commit.Message)

		commitHash = commit.ParentHash
	}
}

func parseCommit(content []byte) *objects.Commit {
	lines := strings.Split(string(content), "\n")
	commit := &objects.Commit{}

	for i, line := range lines {
		if strings.HasPrefix(line, "tree ") {
			commit.TreeHash = line[5:]
		} else if strings.HasPrefix(line, "parent ") {
			commit.ParentHash = line[7:]
		} else if strings.HasPrefix(line, "author ") {
			parts := strings.Split(line[7:], " ")

			if len(parts) >= 2 {
				commit.Author = strings.Join(parts[:len(parts)-1], " ")

				var timestamp int64
				fmt.Sscanf(parts[len(parts)-1], "%d", &timestamp)
				commit.Timestamp = time.Unix(timestamp, 0)
			}
		} else if line == "" && i+1 < len(lines) {
			commit.Message = strings.Join(lines[i+1:], "\n")
			break
		}
	}

	return commit
}

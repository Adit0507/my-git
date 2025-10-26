package commands

import (
	"fmt"

	"github.com/Adit0507/my-git.git/internal/repository"
)

func Status(args []string) {
	index := repository.NewIndex(".")

	if err := index.Load(); err != nil {
		fmt.Printf("error loading index: %v\n", err)
		return
	}

	fmt.Println("Changes to be committed:")
	if len(index.Entries) == 0 {
		fmt.Println("(no files staged)")
	} else {
		for _, entry := range index.Entries {
			fmt.Printf("  new file: %s\n", entry.Path)
		}
	}
}

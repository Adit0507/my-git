package commands

import (
	"fmt"

	"github.com/Adit0507/my-git.git/internal/repository"
)

func Init(args []string) {
	repo := repository.NewRepository(".")
	if err := repo.Init(); err != nil {
		fmt.Printf("Error initializing repo: %v\n", err)
		return
	}

	fmt.Println("Initialized empty VCS repo in .vcs/")
}

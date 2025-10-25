package repository

import "path/filepath"

type Repository struct {
	Path    string
	VcsDir  string
	Storage interface{}
}

func NewRepository(path string) *Repository {
	return &Repository{
		Path:   path,
		VcsDir: filepath.Join(path, ".vcs"),
	}
}


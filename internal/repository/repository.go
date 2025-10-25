package repository

import (
	"os"
	"path/filepath"
)

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

func (r *Repository) Init() error {
	dirs := []string{
		r.VcsDir,
		filepath.Join(r.VcsDir, "objects"),
		filepath.Join(r.VcsDir, "refs"),
		filepath.Join(r.VcsDir, "refs", "heads"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// creatin HEAD file
	headPath := filepath.Join(r.VcsDir, "HEAD")
	if err := os.WriteFile(headPath, []byte("ref: refs/heads/main \n"), 0644); err != nil {
		return err
	}

	return nil
}

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

func (r *Repository) GetHEAD() (string, error) {
	headPath := filepath.Join(r.VcsDir, "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", err
	}

	// parsing ref: refs/head/main
	content := string(data)
	if len(content) > 5 && content[:5] == "ref: " {
		refPath := filepath.Join(r.VcsDir, content[5:len(content)-1])
		refData, err := os.ReadFile(refPath)
		if err != nil {
			if os.IsNotExist(err) {
				return "", nil //no commits
			}

			return "", err
		}

		return string(refData[:len(refData)-1]), nil
	}

	return string(data[:len(data)-1]), nil
}

func (r *Repository) UpdateHEAD(commitHash string) error {
	headPath := filepath.Join(r.VcsDir, "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return err
	}

	content := string(data)
	if len(content) > 5 && content[:5] == "ref: " {
		refPath := filepath.Join(r.VcsDir, content[5:len(content)-1])
		return os.WriteFile(refPath, []byte(commitHash+"\n"), 0644)
	}

	return os.WriteFile(headPath, []byte(commitHash+"\n"), 0644)
}

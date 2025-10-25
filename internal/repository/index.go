package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type IndexEntry struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
	Mode string `json:"mode"`
}

type Index struct {
	Entries map[string]IndexEntry `json:"entries"`
	path    string
}

func NewIndex(repoPath string) *Index {
	return &Index{
		Entries: make(map[string]IndexEntry),
		path:    filepath.Join(repoPath, ".vcs", "index"),
	}
}

func (idx *Index) Load() error {
	data, err := os.ReadFile(idx.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return json.Unmarshal(data, idx)
}

func (idx *Index) Save() error {
	data, err := json.MarshalIndent(idx, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(idx.path, data, 0644)
}

func (idx *Index) Add(path, hash, mode string) {
	idx.Entries[path] = IndexEntry{
		Path: path,
		Hash: hash,
		Mode: mode,
	}
}

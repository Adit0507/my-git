package repository

import "path/filepath"

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

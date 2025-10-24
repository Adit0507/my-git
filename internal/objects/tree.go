package objects

import (
	"bytes"
	"fmt"
	"sort"
)

type TreeEntry struct {
	Mode string
	Name string
	Hash string
}

type Tree struct {
	Entries []TreeEntry
}

func NewTree() *Tree {
	return &Tree{Entries: []TreeEntry{}}
}

func (t *Tree) AddEntry(mode, name, hash string) {
	t.Entries = append(t.Entries, TreeEntry{
		Mode: mode,
		Name: name,
		Hash: hash,
	})
}

func (t *Tree) Type() ObjectType {
	return TreeType
}

func (t *Tree) Serialize() []byte {
	sort.Slice(t.Entries, func(i, j int) bool {
		return t.Entries[i].Name < t.Entries[j].Name
	})

	var buf bytes.Buffer
	for _, entry := range t.Entries{
		buf.WriteString(fmt.Sprintf("%s %s\x00%s\n", entry.Mode, entry.Name, entry.Hash))
	}

	return buf.Bytes()
}

func (t *Tree) Hash()string {
	return  HashContent(TreeType, t.Serialize())
}
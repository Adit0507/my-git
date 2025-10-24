package objects

import (
	"bytes"
	"fmt"
	"time"
)

type Commit struct {
	TreeHash   string
	ParentHash string
	Author     string
	Message    string
	Timestamp  time.Time
}

func NewCommit(treeHash, parentHash, author, message string) *Commit {
	return &Commit{
		TreeHash: treeHash,
		ParentHash: parentHash,
		Author: author,
		Message: message,
		Timestamp: time.Now(),
	}	
}

func (c *Commit) Type() ObjectType {
	return CommitType
}

func (c *Commit) Serialize() []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("tree %s\n", c.TreeHash))
	if c.ParentHash != ""{
		buf.WriteString(fmt.Sprintf("parent %s\n", c.ParentHash))
	}
	buf.WriteString(fmt.Sprintf("author %s %d\n", c.Author, c.Timestamp.Unix()))
	buf.WriteString(fmt.Sprintf("\n%s\n", c.Message))

	return buf.Bytes()
}

func (c *Commit) Hash() string{
	return HashContent(CommitType, c.Serialize())
}
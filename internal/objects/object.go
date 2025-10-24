package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type ObjectType string

const (
	BlobType   ObjectType = "blob" //file content
	TreeType   ObjectType = "tree"
	CommitType ObjectType = "commit"
)

type Object interface {
	Type() ObjectType
	Serialize() []byte
	Hash() string
}

func HashContent(objType ObjectType, content []byte) string {
	header := fmt.Sprintf("%s %d\x00", objType, len(content))
	data := append([]byte(header), content...)
	hash := sha1.Sum(data)

	return hex.EncodeToString(hash[:])
}
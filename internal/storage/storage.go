package storage

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Adit0507/my-git.git/internal/objects"
)

type Storage struct {
	objectsDir string
}

func NewStorage(repoPath string) *Storage {
	return &Storage{
		objectsDir: filepath.Join(repoPath, ".vcs", "objects"),
	}
}

func (s *Storage) getObjectPath(hash string) string {
	return filepath.Join(s.objectsDir, hash[:2], hash[2:])
}

func (s *Storage) WriteObject(obj objects.Object) (string, error) {
	hash := obj.Hash()
	objectPath := s.getObjectPath(hash)

	// checkin if object already exists
	if _, err := os.Stat(objectPath); err != nil {
		return hash, nil
	}

	if err := os.MkdirAll(filepath.Dir(objectPath), 0755); err != nil {
		return "", err
	}

	// wrie compressed object
	file, err := os.Create(objectPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := zlib.NewWriter(file)
	defer writer.Close()

	header := fmt.Sprintf("%s %d\x00", obj.Type(), len(obj.Serialize()))
	if _, err := writer.Write([]byte(header)); err != nil {
		return  "", err
	}
	if _, err := writer.Write(obj.Serialize()); err != nil {
		return "", err
	}

	return hash, nil
}

func (s *Storage) ReadObject(hash string) ([]byte, objects.ObjectType, error) {
	objectPath := s.getObjectPath(hash)

	file, err := os.Open(objectPath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	reader, err := zlib.NewReader(file)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}

	nullIdx := 0
	for i, b := range data {
		if b ==0{
			nullIdx = i
			break
		}
	}

	header := string(data[:nullIdx])

	var objType string
	var size int
	fmt.Sscanf(header, "%s %d", &objType, &size)

	content := data[nullIdx+ 1:]

	return content, objects.ObjectType(objType), nil
}
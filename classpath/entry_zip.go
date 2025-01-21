package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == className {
			classReader, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer classReader.Close()
			data, err := io.ReadAll(classReader)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}

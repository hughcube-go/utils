package msfile

import (
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}

	return s.IsDir()
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}

	return !s.IsDir()
}

func Size(path string) int64 {
	s, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return 0
	}

	if !s.IsDir() {
		return s.Size()
	}

	var size int64
	_ = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size
}

func Mkdir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func Remove(path string) error {
	return os.RemoveAll(path)
}

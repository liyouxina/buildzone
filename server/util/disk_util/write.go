package disk_util

import (
	"io/ioutil"
	"os"
	"strings"
)

func WriteOver(path string, data []byte) error {
	if err := EnsurePathExist(getDirFromFullPath(path)); err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

func EnsurePathExist(path string) error {
	exists, _ := Exists(path)
	if !exists {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

func getDirFromFullPath(path string) string {
	words := strings.Split(path, `/`)
	return strings.Join(words[0: len(words) - 1], `/`)
}
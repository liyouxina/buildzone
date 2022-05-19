package disk

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

func WriteAppend(path string, data []byte) error {
	if err := EnsurePathExist(getDirFromFullPath(path)); err != nil {
		return err
	}
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Errorf("write append error %v", err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	return err
}

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
	return strings.Join(words[0:len(words)-1], `/`)
}

package disk

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func Read(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("error reading file %s", path)
		return nil, err
	}
	return bytes, nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

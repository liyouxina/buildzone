package disk

import (
	"os"
	"testing"
)

const (
	FILE_FOR_TEST_READ      = "/tmp/buildzone/unittest/read_test"
	FILE_FOR_TEST_READ_PATH = "/tmp/buildzone/unittest"
)

func TestExists(t *testing.T) {
	err := EnsurePathExist(FILE_FOR_TEST_READ_PATH)
	if err != nil {
		t.Fail()
	}
	err = os.WriteFile(FILE_FOR_TEST_READ, []byte{'a'}, os.ModePerm)
	defer os.Remove(FILE_FOR_TEST_READ)
	if err != nil {
		t.Fail()
	}
	exists, err := Exists(FILE_FOR_TEST_READ)
	if err != nil {
		t.Fail()
	}
	if !exists {
		t.Fail()
	}
}

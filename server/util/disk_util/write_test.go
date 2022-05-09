package disk_util

import "testing"

const (
	FILE_FOR_TEST = "/tmp/buildzone/test_write_over"
	FILE_DIR_FOR_TEST_= "/tmp/buildzone"
)

func Test_getDirFromFullPath(t *testing.T)  {
	if (getDirFromFullPath(FILE_FOR_TEST) != FILE_DIR_FOR_TEST_) {
		t.Fail()
	}
}

func TestEnsurePathExist(t *testing.T) {
	if err := EnsurePathExist(FILE_DIR_FOR_TEST_); err != nil {
		t.Errorf("error EnsurePathExist %v", err)
		t.Fail()
	}
}

func TestWriteOver(t *testing.T) {
	content := []byte{'a', 'b', 'c'}
	err := WriteOver(FILE_FOR_TEST, content)
	if err != nil {
		t.Errorf("err writeover %v", err)
		t.Fail()
	}
	readContent, err := Read(FILE_FOR_TEST)
	if err != nil {
		t.Errorf("err read %v", err)
		t.Fail()
	}
	equal := byteArrayEqual(readContent, content)
	if !equal {
		t.Log("write not right")
		t.Fail()
	}
	err = WriteOver(FILE_FOR_TEST, content)
	if err != nil {
		t.Errorf("err writeover %v", err)
		t.Fail()
	}
	readContent, err = Read(FILE_FOR_TEST)
	if err != nil {
		t.Errorf("err read %v", err)
		t.Fail()
	}
	equal = byteArrayEqual(readContent, content)
	if !equal {
		t.Log("write not right")
		t.Fail()
	}
}

func byteArrayEqual(bytes1 []byte, bytes2 []byte) bool {
	if len(bytes1) != len(bytes2) {
		return false
	}
	for index, b := range bytes1 {
		if bytes2[index] != b {
			return false
		}
	}
	return true
}

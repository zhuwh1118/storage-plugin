package main

import (
	"testing"
)

func TestFile(t *testing.T) {
	println("start test")
	s2 := New(map[string]string{
		"region":          "",
		"bucketname":      "",
		"accessKeyId":     "",
		"accessKeySecret": "",
	})

	f1 := s2.File("filename1")     //truely exsit
	f2 := s2.File("filename2")     // turely not exsit
	f3 := s2.File("just for test") // turely not exsit
	if f1.Key() != "filename1" {
		t.Error("err key")
	}
	result, _, err1 := f1.Exist()
	if err1 != nil || result != true {
		t.Error("err exist")
	}
	result, _, err2 := f2.Exist()
	if err2 != nil || result != false {
		t.Error("err exsit2")
	}
	bt, _, err := f1.Bytes()
	if len(bt) == 0 || err != nil {
		t.Error("err bytes")
	}
	_, err = f2.Delete()
	if err != nil {
		t.Error("err delete")
	}
	length, _, err := f3.Append([]byte("2018-0521"), 0)
	if err != nil || int(length) != len([]byte("2018-0521")) {
		t.Error("err append")
	}
	length, _, err = f3.Append([]byte("a"), 9)
	if err != nil || int(length) != len([]byte("2018-0521a")) {
		t.Error("err append1", int(length))
	}
}

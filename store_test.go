package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "0e079099d03aca5482d88be7e4b32b952a983c62"
	expectedPathName := "0e079/099d0/3aca5/482d8/8be7e/4b32b/952a9/83c62"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Original != expectedPathName {
		t.Errorf("have %s want %s", pathKey.Original, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "myspecialpic"
	pathKey := CASPathTransformFunc(key)
	expectedPathname := "87438/fc7f5/7665f/69661/2bd71/7904a/fc598/1854d"
	expectedOriginalKey := "87438fc7f57665f696612bd717904afc5981854d"

	if pathKey.Pathname != expectedPathname {
		t.Errorf("have %s want %s", pathKey.Pathname, expectedPathname)
	}

	if pathKey.Original != expectedOriginalKey {
		t.Errorf("have %s want %s", pathKey.Original, expectedOriginalKey)
	}	
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))

	if err := s.writeStream("myspecialpic", data); err != nil {
		t.Error(err)
	}
}
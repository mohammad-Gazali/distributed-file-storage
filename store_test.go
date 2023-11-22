package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "myspecialpic"
	pathKey := CASPathTransformFunc(key)
	expectedPathname := "87438/fc7f5/7665f/69661/2bd71/7904a/fc598/1854d"
	expectedFilename := "87438fc7f57665f696612bd717904afc5981854d"

	if pathKey.Pathname != expectedPathname {
		t.Errorf("have %s want %s", pathKey.Pathname, expectedPathname)
	}

	if pathKey.Filename != expectedFilename {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFilename)
	}	
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	key := "supersecretkey"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	key := "supersecretkey"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)

	if err != nil {
		t.Error(err)
	}

	b, err := io.ReadAll(r)

	if err != nil {
		t.Error(err)
	}

	if string(b) != string(data) {
		t.Errorf("have %s want %s", b, data)
	}

	s.Delete(key)
}

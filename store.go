package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

type PathTransformFunc func(string) PathKey

func DefaultPathTransformFunc(key string) PathKey {
	return PathKey{
		Pathname: key,
		Filename: key,
	}
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i * blocksize, i * blocksize + blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Filename: hashStr,
	}
}

type StoreOpts struct {
	PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Has(key string) bool {
	pathKey := s.PathTransformFunc(key)

	_, err := os.Stat(pathKey.FullPath())
	
	return err != fs.ErrNotExist
}

func (s *Store) Delete(key string) error {
	var err error
	pathKey := s.PathTransformFunc(key)

	defer func() {
		if err == nil {
			log.Printf("deleted [%s] from disk successfully", pathKey.Filename)
		}
	}()

	err = os.Remove(pathKey.FullPath())
	return err
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, f)

	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	return os.Open(pathKey.FullPath())
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	pathAndFilename := pathKey.FullPath()

	f, err := os.Create(pathAndFilename)

	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)

	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s\n", n, pathAndFilename)

	return nil
}

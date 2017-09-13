package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/datadog/zstd"
)

// ZSTDStore implements the storage of files compressed with zstd
type ZSTDStore struct{}

// Store stores data in compressed form in a file
func (z ZSTDStore) Store(path string, data []byte) error {
	d, err := zstd.Compress(nil, data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(gzPath(path), d, 0644)
}

// Load returns uncompressed data from a compressed file
func (z ZSTDStore) Load(path string) ([]byte, error) {
	d, err := ioutil.ReadFile(gzPath(path))
	if err != nil {
		return ioutil.ReadFile(strings.TrimSuffix(path, ".gz"))
	}
	return zstd.Decompress(nil, d)
}

// Compress existing file
func (z ZSTDStore) Compress(path string) error {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		log.Println(err)
	}

	return z.Store(path, d)
}

// Decompress existing file
func (z ZSTDStore) Decompress(path string) error {
	d, err := z.Load(path)
	if err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		log.Println(err)
	}

	return ioutil.WriteFile(strings.TrimSuffix(path, ".gz"), d, 0644)
}

func gzPath(path string) string {
	if path[len(path)-3:] != ".gz" {
		path += ".gz"
	}
	return path
}

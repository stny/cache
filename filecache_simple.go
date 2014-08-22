package cache

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"sync"
)

type FileCacheSimple struct {
	mu  sync.Mutex
	dir string
}

func NewFileCacheSimple(d string) *FileCacheSimple {
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.Mkdir(d, 0755); err != nil {
			panic(err)
		}
	}
	return &FileCacheSimple{dir: d}
}

func (c *FileCacheSimple) Write(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	filename := generateFilename(key)
	if err := atomicWriteFile(filename, c.dir, []byte(value)); err != nil {
		return err
	}
	return nil
}

func (c *FileCacheSimple) Read(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	filename := generateFilename(key)
	buf, err := ioutil.ReadFile(path.Join(c.dir, filename))
	if err != nil {
		return "", false
	}
	return string(buf), true
}

func generateFilename(key string) string {
	key = url.QueryEscape(key)
	return key
}

func atomicWriteFile(filename string, dir string, data []byte) error {
	tempFile, err := ioutil.TempFile(dir, filename)
	if err != nil {
		return err
	}

	defer func() {
		os.Remove(tempFile.Name())
		tempFile.Close()
	}()

	tempFile.Write(data)
	if err := tempFile.Sync(); err != nil {
		return err
	}

	if err := os.Rename(tempFile.Name(), path.Join(dir, filename)); err != nil {
		return err
	}
	return nil
}

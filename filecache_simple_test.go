package cache

import "testing"

func TestFileSimpleReadAndWrite(t *testing.T) {
	c, _ := NewFileCacheSimple("./tmp")
	readAndWrite(t, c)
}

func BenchmarkFileSimple(b *testing.B) {
	c, _ := NewFileCacheSimple("./tmp")
	readAndWriteForBench(b, c)
}

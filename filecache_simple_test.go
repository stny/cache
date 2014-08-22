package cache

import "testing"

func TestFileSimpleReadAndWrite(t *testing.T) {
	c := NewFileCacheSimple("./tmp")
	readAndWrite(t, c)
}

func BenchmarkFileSimple(b *testing.B) {
	c := NewFileCacheSimple("./tmp")
	readAndWriteForBench(b, c)
}

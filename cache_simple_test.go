package cache

import "testing"

func TestSimpleReadAndWrite(t *testing.T) {
	c := NewCacheSimple()
	readAndWrite(t, c)
}

func BenchmarkSimple(b *testing.B) {
	c := NewCacheSimple()
	readAndWriteForBench(b, c)
}

package cache

import "testing"

const (
	key   = "harapeko"
	value = "(╯°□°）╯︵ ┻━┻"
)

type cacher interface {
	Read(string) (string, bool)
	Write(string, string) error
}

func readAndWrite(t *testing.T, c cacher) {
	if err := c.Write(key, value); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if got, ok := c.Read(key); !ok || value != got {
		t.Fatalf("%#v != %#v", value, got)
	}
}

func readAndWriteForBench(b *testing.B, c cacher) {
	c.Write(key, value)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, ok := c.Read(key)
		if !ok {
			b.Fatalf("Unexpected error")
		}
	}
}

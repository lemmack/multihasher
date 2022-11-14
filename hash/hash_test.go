package hash

import (
	"testing"
)

func BenchmarkMakeHashJson(b *testing.B) {
	bytes := []byte("this is a test string")

	for i := 0; i < b.N; i++ {
		_, _ = MakeHashJson(bytes)
	}
}

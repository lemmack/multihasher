package server

import (
	"testing"
)

func BenchmarkServer(b *testing.B) {
	bytes := []byte("this is a test string")

	for i := 0; i < b.N; i++ {
		_, _ = make_hash_json(bytes)
	}
}

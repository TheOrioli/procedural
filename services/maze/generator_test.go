package maze

import (
	"math/rand"
	"testing"
)

func BenchmarkGenerateBacktrack10x10(b *testing.B) {
	src := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		generate(10, 10, src, backtrack{})
	}
}

func BenchmarkGenerateBacktrack100x100(b *testing.B) {
	src := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		generate(100, 100, src, backtrack{})
	}
}

func BenchmarkGenerateBacktrack1000x1000(b *testing.B) {
	src := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		generate(1000, 1000, src, backtrack{})
	}
}

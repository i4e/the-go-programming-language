package popcount

import (
	"math/rand"
	"testing"
)

func TestPopCount(t *testing.T) {
	for i := uint64(0); i < 1000000; i++ {
		pc := PopCount(i)
		if got := PopCountLoop(i); got != pc {
			t.Errorf("want: %v, got: %v", pc, got)
		}
		if got := PopCountShift(i); got != pc {
			t.Errorf("want: %v, got: %v", pc, got)
		}
		if got := PopCountClear(i); got != pc {
			t.Errorf("want: %v, got: %v", pc, got)
		}
	}
}

func benchmarkPopCount(b *testing.B, f func(uint64) int) {
	b.Helper()

	rand.Seed(0)
	s := 0
	for i := 0; i < b.N; i++ {
		x := rand.Uint64()
		s += f(x)
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmarkPopCount(b, PopCount)
}

func BenchmarkPopCountLoop(b *testing.B) {
	benchmarkPopCount(b, PopCountLoop)
}

func BenchmarkPopCountShift(b *testing.B) {
	benchmarkPopCount(b, PopCountShift)
}

func BenchmarkPopCountClear(b *testing.B) {
	benchmarkPopCount(b, PopCountClear)
}

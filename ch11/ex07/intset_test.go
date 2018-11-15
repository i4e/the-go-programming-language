package intset

import (
	"math/rand"
	"testing"
)

const maxnum = 100000000
const setSize = 100000

func BenchmarkIntSetAdd(b *testing.B) {
	rand.Seed(0)

	for i := 0; i < b.N; i++ {
		is := &IntSet{}
		for j := 0; j < setSize; j++ {
			is.Add(rand.Intn(maxnum))
		}
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	rand.Seed(0)
	isa := &IntSet{}
	for j := 0; j < setSize; j++ {
		isa.Add(rand.Intn(maxnum))
	}
	isb := &IntSet{}
	for j := 0; j < setSize; j++ {
		isb.Add(rand.Intn(maxnum))
	}

	for i := 0; i < b.N; i++ {
		isa.UnionWith(isb)
	}
}

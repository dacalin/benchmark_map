package main

import (
	"github.com/dolthub/swiss"
	"testing"
)

func BenchmarkGoMapString(b *testing.B) {
	m := make(map[int]string, 1000)
	for i := 0; i < b.N; i++ {
		m[i] = ""
	}
}

func BenchmarkSwissMapString(b *testing.B) {
	m := swiss.NewMap[int, string](1000)
	for i := 0; i < b.N; i++ {
		m.Put(i, "")
	}
}

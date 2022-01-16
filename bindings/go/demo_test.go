package demo

import (
	"testing"
)

func TestDemoCAdd(t *testing.T) {
	if Add(1, 3) != 4 {
		t.Fatalf(`TestDemoAdd Failed`)
	}
}

func BenchmarkDemoCAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CAdd(1111, 1111)
	}
}

func BenchmarkDemoAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1111, 1111)
	}
}

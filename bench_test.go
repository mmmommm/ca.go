package main

import (
	"testing"
)

// Functional Option Pattern (FOP)
func BenchmarkFunctionalOption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewBicycleWithFOP(Road,
        WithBodyType("almi"),
        WithTotalGear(4, 8),
		)
	}
}

// Builder Pattern (BP)
func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewBicycleWithBP(Mountain).
        WithBodyType("Carbon").
        WithTotalGear(3, 9).
        Build()
	}
}

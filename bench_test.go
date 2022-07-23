package main

import (
	"testing"
)

// Functional Option Pattern (FOP)
func BenchmarkFunctionalOption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewApplicationWithFOP(Premium,
			WithBackupService(true),
			WithSupport(true),
			WithMovie(false),
		)
	}
}

// Builder Pattern (BP)
func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewApplicationWithBP(Premium).
			WithBackupService(true).
			WithSupport(true).
			WithMovie(false).
			Build()
	}
}

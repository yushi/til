package main

import "testing"

func BenchmarkNormalAppend(b *testing.B) {
	s := []int{}
	for i := 0; i < b.N; i++ {
		s = append(s, 1)
	}
}
func BenchmarkIndex(b *testing.B) {
	s := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		s[i] = 1
	}
}
func BenchmarkPreallocatedAppend(b *testing.B) {
	s := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, 1)
	}
}

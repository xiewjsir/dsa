package main

import (
	"strings"
	"testing"
)

var (
	str = "abababbcabcac"
	p   = "abc"
)

func VM() {
	ViolenceMatch(str, p)
}

//kmp 算法
func KmpSearch() {
	Kmp(str, p, GetNext(p))
}

//官方使用的 Rabin-Karp search
func RKSearch() {
	strings.Index(str, p)
}

func BenchmarkKmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KmpSearch()
	}
}

func BenchmarkPK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RKSearch()
	}
}

func BenchmarkVM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VM()
	}
}

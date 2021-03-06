package gotesting

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	// test itself
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	// test itself
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	// test itself
	for i := 0; i < b.N; i++ {
		h := sha512.New() // takes 50% longer
		h.Sum(data)       // allocating memory
	}
}

// go test -bench
// runs bench tests and unit tests
// for the module

// go test -bench -benchtime 10s

// go test -bench .-benchtime 10s
// <- end-to-end -benchtime must be higher seconds
// 99 million interations to reach our ten second target

// go tool pprof profile.out
// svg <- generates a svg type file for profile info
// graphviz

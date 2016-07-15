package escapist

import (
	"fmt"
	"html"
	"testing"
)

var (
	exampleBasicStr      = "Hello there Mr. Joe <inject stuff>"
	exampleBasicNoRepStr = "Hello there Mr. Joe. This is a safe string"

	exampleBasic      = []byte(exampleBasicStr)
	exampleBasicNoRep = []byte(exampleBasicNoRepStr)

	outB   []byte
	outStr string
)

func TestBasic(t *testing.T) {
	b := Escape(exampleBasic)
	fmt.Println(string(b))
}

func TestBasicNoRep(t *testing.T) {
	b := Escape(exampleBasicNoRep)
	fmt.Println(string(b))
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = Escape(exampleBasic)
	}

	b.ReportAllocs()
}

func BenchmarkNoRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = Escape(exampleBasicNoRep)
	}

	b.ReportAllocs()
}

func BenchmarkHTMLBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outStr = html.EscapeString(exampleBasicStr)
	}

	b.ReportAllocs()
}

func BenchmarkHTMLNoRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outStr = html.EscapeString(exampleBasicNoRepStr)
	}

	b.ReportAllocs()
}

func BenchmarkAdvBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = EscapeAdv(exampleBasic)
	}

	b.ReportAllocs()
}

func BenchmarkAdvNoRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = EscapeAdv(exampleBasicNoRep)
	}

	b.ReportAllocs()
}

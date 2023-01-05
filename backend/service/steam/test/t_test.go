package test

import (
	"testing"
)

//
//func TestName1(t *testing.T) {
//
//}
//
//func TestName2(t *testing.T) {
//
//}

type t1 struct {
	Id   int
	Name string
}

func (t t1) Create() {
	// do something
	t.Name = "from create"
}

func a1() {
	t1{
		Id: 0,
	}.Create()
}

func a2() {
	t := t1{
		Id: 0,
	}
	t.Create()
}

func BenchmarkName1(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		a1()
	}
}

func BenchmarkName2(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		a2()
	}
}

package main

import "testing"

// Exercise 1.3:
// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join

var input1 = []string{"a", "b", "c", "d"}
var input2 = []string{"a", "b", "c", "d", "arg5", "arg6", "arg6", "arg7"}
var input3 = []string{
	"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz1", "234",
	"567", "890", "aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh",
	"iii", "jjj", "kkk", "lll", "mmm", "nnn", "ooo", "ppp", "qqq", "rrr",
	"sss", "ttt", "uuu", "vvv", "www", "xxx", "yyy", "zzz", "abc1", "def2",
	"ghi3", "jkl4", "mno5", "pqr6", "stu7", "vwx8", "yz19", "2340", "5671", "8902",
	"a1a", "b2b", "c3c", "d4d", "e5e", "f6f", "g7g", "h8h", "i9i", "j0j",
	"kak", "lbl", "lcl", "mdm", "nen", "ofo", "pgp", "qhq", "iri", "sjs",
	"tkt", "ulu", "vmv", "wnw", "xox", "ypy", "zqz", "a1b", "b2c", "c3d",
	"d4e", "e5f", "f6g", "g7h", "h8i", "i9j", "j0k", "k1l", "l2m", "m3n",
	"n4o", "o5p", "p6q", "q7r", "r8s", "s9t", "t0u", "u1v", "v2w", "w3x",
}

var input = input3

// for loop
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(input)
	}
}

// for range
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2(input)
	}
}

// strings.Join
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(input)
	}
}

// strings.Builder
func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo4(input)
	}
}

package main

import "testing"

func Test(t *testing.T) {
	rez1, rez2 := NumPassingPasswords("TestFile.txt")
	exp1 := 2
	exp2 := 1
	if rez1 != exp1 || rez2 != exp2 {
		t.Errorf("Password count did not match")
	}
}

func Test2(t *testing.T) {
	rez1, rez2 := NumPassingPasswords("RealFile.txt")
	exp1 := 666
	exp2 := 670
	if rez1 != exp1 || rez2 != exp2 {
		t.Errorf("Password count did not match")
	}
}

func BenchmarkNumPassing(b *testing.B) {
	exp1 := 666
	exp2 := 670
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rez1, rez2 := NumPassingPasswords("RealFile.txt")
		if rez1 != exp1 || rez2 != exp2 {
			b.Errorf("Password count did not match")
		}
	}
}

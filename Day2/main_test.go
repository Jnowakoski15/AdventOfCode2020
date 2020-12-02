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

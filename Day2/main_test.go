package main

import "testing"

func Test(t *testing.T) {
	rez := NumPassingPasswords("TestFile.txt")
	exp := 2

	if rez != exp {
		t.Errorf("Password count did not match")
	}
}

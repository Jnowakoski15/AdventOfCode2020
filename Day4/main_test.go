package main

import "testing"

const testName = "testFile.txt"

func TestPart1Example(t *testing.T) {
	exp := 2
	rez := part1Evaluation(testName)

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestPart2ValidExample(t *testing.T) {
	exp := 4
	rez := part2Evaluation("testValidFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestBirthYearInBounds(t *testing.T) {
	in := "1991"
	exp := true
	rez := birthYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestBirthYearUnderBounds(t *testing.T) {
	in := "0"
	exp := false
	rez := birthYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestBirthYearAboveBounds(t *testing.T) {
	in := "99999"
	exp := false
	rez := birthYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestIssueYearInBounds(t *testing.T) {
	in := "2015"
	exp := true
	rez := issueYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestIssueYearUnderBounds(t *testing.T) {
	in := "2000"
	exp := false
	rez := issueYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestIssueYearAboveBounds(t *testing.T) {
	in := "2025"
	exp := false
	rez := issueYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestExperationYearTestInBounds(t *testing.T) {
	in := "2025"
	exp := true
	rez := experationYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestExperationYearTestUnderBounds(t *testing.T) {
	in := "2019"
	exp := false
	rez := experationYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestExperationYearTestAboveBounds(t *testing.T) {
	in := "2031"
	exp := false
	rez := experationYearTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestEyeColorForTrueResponse(t *testing.T) {
	in := "amb"
	exp := true
	rez := eyeColorTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestEyeColorForFalse(t *testing.T) {
	in := "potato"
	exp := false
	rez := eyeColorTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHairColorTestForTrue(t *testing.T) {
	in := "#888785"
	exp := true
	rez := hairColorTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHairColorTestForTooFewChars(t *testing.T) {
	in := "#FFF"
	exp := false
	rez := hairColorTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHairColorTestLackingPrefix(t *testing.T) {
	in := "888785"
	exp := false
	rez := hairColorTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestPassportIDForTrueResponse(t *testing.T) {
	in := "123456789"
	exp := true
	rez := passportIDTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestPassportIDForTooManyDigits(t *testing.T) {
	in := "1234567890"
	exp := false
	rez := passportIDTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestPassportIDForTooFewDigits(t *testing.T) {
	in := "1234"
	exp := false
	rez := passportIDTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightCMInbounds(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "150cm"
	exp := true
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightCMUnderBounds(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "100cm"
	exp := false
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightCMAboveBounds(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "200cm"
	exp := false
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightINInbounds(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "60in"
	exp := true
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightinUnderBounds(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "58in"
	exp := false
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightINAboveBounds(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "200in"
	exp := false
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestHeightLackingSuffix(t *testing.T) {
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	in := "200"
	exp := false
	rez := heightTest(in)
	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

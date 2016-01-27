package letterDiamond

import "testing"

func TestIsValidInput(t *testing.T) {
    if IsValidInput("A") != true {
        t.Errorf("IsValidInput(%q) == false, expected true", "A")
    }
    if IsValidInput("*") != false {
        t.Errorf("IsValidInput(%q) == true, expected false", "*")
    }
}

func TestLetterString(t *testing.T) {
    cases := []struct {
		letter, letters string
	}{
		{"A", "A"},
		{"F", "ABCDEF"},
		{"K", "ABCDEFGHIJK"},
	}
	for _, c := range cases {
        letters := LetterString(c.letter)
        if letters != c.letters {
            t.Errorf("LetterString(%q) == %q, expected %q", c.letter, letters, c.letters)
        }
    }
}


func TestStringForLetter(t *testing.T) {
    abc, a, b, c  := "ABC", "A", "B", "C"
    aString, aStringExpected := StringForLetter(a, 0, 3), "  A  "
    bString, bStringExpected := StringForLetter(b, 1, 3), " B B "
    cString, cStringExpected := StringForLetter(c, 2, 3), "C   C"

    if aString != aStringExpected {
        t.Errorf("StringForLetter(%q, %q) == %q, expected %q", a, abc, aString, aStringExpected)
    }
    
    if bString != bStringExpected {
        t.Errorf("StringForLetter(%q, %q) == %q, expected %q", b, abc, bString, bStringExpected)
    }
    
    if cString != cStringExpected {
        t.Errorf("StringForLetter(%q, %q) == %q, expected %q", c, abc, cString, cStringExpected)
    }
}

func TestLetterStrings(t *testing.T) {
    aStrings := LetterStrings("A")
    len_a := len(aStrings)
    if len_a != 1 {
        t.Errorf("len(LetterStrings(%q)) == %d, expected %d", "A", len_a, 1)
    }

    if aStrings[0] != "A" {
        t.Errorf("LetterStrings(%q)[0] == %q, expected %q", "A", aStrings[0], "A")
    }
    
    dStrings := LetterStrings("D")
    len_d := len(dStrings)
    if len_d != 4 {
        t.Errorf("len(LetterStrings(%q)) == %d, expected %d", "D", len_d, 4)
    }

    if dStrings[0] != "   A   " {
        t.Errorf("LetterStrings(%q)[0] == %q, expected %q", "D", dStrings[0], "   A   ")
    }

    if dStrings[1] != "  B B  " {
        t.Errorf("LetterStrings(%q)[0] == %q, expected %q", "B", dStrings[1], "B     B")
    }

    if dStrings[2] != " C   C " {
        t.Errorf("LetterStrings(%q)[0] == %q, expected %q", "C", dStrings[2], "C     C")
    }
    
    if dStrings[3] != "D     D" {
        t.Errorf("LetterStrings(%q)[0] == %q, expected %q", "D", dStrings[3], "D     D")
    }
}
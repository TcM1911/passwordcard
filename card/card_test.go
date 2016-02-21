package card

import (
	"bytes"
	"github.com/tcm1911/passwordcard/utils"
	"testing"
)

// Test createHeader function.
func TestcreateHeader(t *testing.T) {
	cases := []struct {
		input, expected []byte
	}{
		{[]byte("abcde"), []byte("eabcd")},
		{[]byte("abcdefgh"), []byte("habcd")},
	}
	prg := new(mockRandNum)

	for _, c := range cases {
		got := createHeader(c.input, 5, prg)
		if bytes.Compare(got, c.expected) != 0 {
			t.Error("Incorrect array returned. Got:", got, "expected:", c.expected)
		}
	}
}

func (rnd *mockRandNum) nextInt(i int) int {
	if i == 1 {
		return 0
	}
	return i - 2
}

type mockRandNum struct {
	utils.Randomer
}

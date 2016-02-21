package card

import (
	"bytes"
	"github.com/tcm1911/passwordcard/utils"
	"testing"
)

// Test createHeader function.
func TestcreateHeader(t *testing.T) {
	headerChars := []byte("abcde")
	expected := []byte("eabcd")
	prg := new(mockRandNum)
	got := createHeader(headerChars, 5, prg)
	if bytes.Compare(got, expected) != 0 {
		t.Error("Incorrect array returned. Got:", got, "expected:", expected)
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

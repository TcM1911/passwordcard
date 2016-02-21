package utils

import "testing"
import "bytes"

func TestSwap(t *testing.T) {
	initial := []byte("abc")
	expected := []byte("bac")
	swap(initial, 0, 1)
	if bytes.Compare(initial, expected) != 0 {
		t.Error("Swap fail!")
	}
}

func TestShufle(t *testing.T) {
	array := []byte("abcde")
	expected := []byte("edcba")
	rnd := new(generator)
	Shufle(array, rnd)
	if bytes.Compare(array, expected) != 0 {
		t.Error("Arrays should be the same: ", array, expected)
	}
}

/*

   Helper functions

*/

// Mock Randomer
type generator struct {
}

func (rnd *generator) nextInt(i int) int {
	var resp int
	switch i {
	case 5:
		resp = 0
	case 4:
		resp = 1
	case 3:
		resp = 2
	case 2:
		resp = 1
	case 1:
		resp = 0
	default:
		panic("Mock error. Invalid input.")
	}
	return resp
}

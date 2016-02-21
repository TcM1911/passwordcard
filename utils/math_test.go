package utils

import "testing"

func TestMathMod(t *testing.T) {
	cases := []struct {
		n, m, expected int
	}{
		{9, 3, 0},
		{0, 1, 0},
		{-9, 3, 0},
		{10, 3, 1},
		{-10, 3, 2},
	}
	for _, c := range cases {
		got := MathMod(c.n, c.m)
		if got != c.expected {
			t.Error("Calling", c.n, "mod", c.m, "expected:", c.expected, "got", got)
		}
	}
}

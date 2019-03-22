package main

import (
	"testing"
)


func Test_MaximalSquare_1(t *testing.T) {
	in := []string{"0111", "1111", "1111", "1111"}
	expected := 9
	actual := MaximalSquare(in)
	if actual != expected {
		t.Error("Actual should be expected")
	}
}

func Test_MaximalSquare_2(t *testing.T) {
	in := []string{"10100", "10111", "11111", "10010"}
	expected := 4
	actual := MaximalSquare(in)
	if actual != expected {
		t.Error("Actual should be expected")
	}
}
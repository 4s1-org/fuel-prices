package main

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(1, 2)
	if result != 3 {
		t.Error("incorrect result: expect 5, got", result)
	}
}

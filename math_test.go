package main

import "testing"

func TestSuM(t *testing.T) {
	total := Sum(15, 15)
	if total != 30 {

		t.Errorf("Wrong result! Result %d. Expected: %d", total, 30)
	}
}

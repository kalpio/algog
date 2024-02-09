package main

import "testing"

func TestInsertionSort(t *testing.T) {
	equal := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}

		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	data := []int{4, 3, 1, 5, 2}
	got := InsertionSort(data)
	exp := []int{1, 2, 3, 4, 5}

	if !equal(got, exp) {
		t.Fail()
	}
}

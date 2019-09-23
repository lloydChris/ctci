package main

import "testing"

func TestCheckUnique(t *testing.T) {
	var tests = make(map[string]bool)
	tests["red"] = true
	tests["house"] = true
	tests["colour"] = false
	tests["!@8*"] = true
	tests["qwertyuiopasdfghjklzxvbnm,./;'[]"] = true

	for text, expected := range tests {
		actual := CheckUnique(text)
		if actual != expected {
			t.Error(text)
		}
	}

}

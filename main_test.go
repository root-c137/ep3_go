package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := add(2, 2)

	result += 1
	if result != 4 {
		t.Fatalf("\nLe résultat de 2+2 doit être égal à 4 et non pas : %v\n", result)
	}
}

package wordlemlwordlists

import (
	"strings"
	"testing"
)

func TestValidGuessesCSV(t *testing.T) {
	lines := strings.Split(ValidGuessesCSV(), "\n")

	if len(lines) != 12947 {
		t.Errorf("Expected 12947 valid guesses, got %d", len(lines))
	}
	if lines[0] != "AAHED" {
		t.Errorf("Expected first valid guess to be AAHED, got %q", lines[0])
	}
}

func TestValidSolutionsCSV(t *testing.T) {
	lines := strings.Split(ValidSolutionsCSV(), "\n")

	if len(lines) != 2309 {
		t.Errorf("Expected 2309 valid solutions, got %d", len(lines))
	}
	if lines[0] != "ABACK" {
		t.Errorf("Expected first valid solution to be ABACK, got %q", lines[0])
	}
}

package wordlemlwordlists

import (
	"strings"
	"testing"
)

func TestValidGuessesCSV(t *testing.T) {
	assertWordlist(t, "valid guesses", ValidGuessesCSV(), 12947, "AAHED", "ZYMIC")
}

func TestValidSolutionsCSV(t *testing.T) {
	assertWordlist(t, "valid solutions", ValidSolutionsCSV(), 2309, "ABACK", "ZONAL")
}

func TestActionSpaceCSV(t *testing.T) {
	assertWordlist(t, "action space", ActionSpaceCSV(), 4739, "AARGH", "ZULUS")
}

func assertWordlist(t *testing.T, name string, csv string, wantCount int, wantFirst string, wantLast string) {
	t.Helper()

	if strings.HasSuffix(csv, "\n") {
		t.Errorf("Expected %s wordlist not to end with a newline", name)
	}

	lines := strings.Split(csv, "\n")

	if len(lines) != wantCount {
		t.Errorf("Expected %d %s words, got %d", wantCount, name, len(lines))
	}
	if lines[0] != wantFirst {
		t.Errorf("Expected first %s word to be %s, got %q", name, wantFirst, lines[0])
	}
	if lines[len(lines)-1] != wantLast {
		t.Errorf("Expected last %s word to be %s, got %q", name, wantLast, lines[len(lines)-1])
	}
}

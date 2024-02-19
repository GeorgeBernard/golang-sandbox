package ordinals

import (
	"fmt"
	"testing"

	"github.com/jaymcgavren/ordinals"
)

type testData struct {
	number   int
	expected string
}

func errorString(test testData, got string) string {
	return fmt.Sprintf("Ordinal(%d) = \"%s\", want \"%s\"", test.number, got, test.expected)
}

func TestOrdinals(t *testing.T) {
	tests := []testData{
		{number: 1, expected: "1st"},
		{number: 2, expected: "2nd"},
		{number: 3, expected: "3rd"},
		{number: 4, expected: "4th"},
		{number: 11, expected: "11th"},
		{number: 21, expected: "21st"},
	}

	for _, test := range tests {
		got := ordinals.Ordinal(test.number)
		if got != test.expected {
			t.Error(errorString(test, got))
		}
	}
}

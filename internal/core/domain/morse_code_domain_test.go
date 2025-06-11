package domain

import (
	"testing"
)

func TestItShouldReturnAnErrorWhenAttemptingToGenerateTheOutput(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		expected string
	}{
		{
			scenario: "Case 1",
			input:    ".... . -.-B-   .-Y-- ..- -.. .A",
			expected: "invalid input(.... . -.-B-   .-Y-- ..- -.. .A). Only Morse code following the standard provided table should be entered",
		},
		{
			scenario: "Case 2",
			input:    "Ç.... . -.--  4 .--- F ..- G -.. . A",
			expected: "invalid input(Ç.... . -.--  4 .--- F ..- G -.. . A). Only Morse code following the standard provided table should be entered",
		},
		{
			scenario: "Case 3",
			input:    " Ç .... . -.--  Z .--- F ..-  G  -.. . A ",
			expected: "invalid input( Ç .... . -.--  Z .--- F ..-  G  -.. . A ). Only Morse code following the standard provided table should be entered",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			morseCode := NewMorseCodeDomain(test.input)
			err := morseCode.Process()
			if err != nil && err.Error() != test.expected {
				t.Fatalf("input: %q\nexpected: %q", test.input, test.expected)
			}

			output := morseCode.GetOutput()
			if len(output) > 0 {
				t.Fatalf("input: %q\nexpected: %q\ngot: %q", test.input, "", output)
			}
		})
	}
}

func TestItShouldGenerateTheOutputWithoutErrors(t *testing.T) {

	tests := []struct {
		scenario string
		input    string
		expected string
	}{
		{
			scenario: "Case 1",
			input:    ".... . -.--   .--- ..- -.. .",
			expected: "HEY JUDE",
		},
		{
			scenario: "Case 2",
			input:    ".... . .-.. .-.. ---   .-- --- .-. .-.. -..",
			expected: "HELLO WORLD",
		},
		{
			scenario: "Case 3",
			input:    ".... . .-.. .-.. ---   .-- --- .-. .-.. -..   .... . -.--   .--- ..- -.. .",
			expected: "HELLO WORLD HEY JUDE",
		},
		{
			scenario: "Case 4",
			input:    " .... . -.--   .--- ..- -.. . ",
			expected: "HEY JUDE",
		},
		{
			scenario: "Case 5",
			input:    "  ....  .  -.--   .---  ..-  -..  .  ",
			expected: "HEY JUDE",
		},
		{
			scenario: "Case 6",
			input:    "  ....  .  -.--",
			expected: "HEY",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			morseCode := NewMorseCodeDomain(test.input)
			_ = morseCode.Process()

			output := morseCode.GetOutput()
			if output != test.expected {
				t.Fatalf("input: %q\nexpected: %q\ngot: %q", test.input, test.expected, output)
			}
		})
	}
}

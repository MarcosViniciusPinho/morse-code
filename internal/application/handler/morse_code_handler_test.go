package handler

import (
	"testing"

	"github.com/MarcosViniciusPinho/morse-code/internal/application/dto"
)

func TestItShouldReturnAnErrorWhenAttemptingToGenerateTheOutputEncapsulatedInTheDto(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		expected string
	}{
		{
			scenario: "Case 1",
			input:    ".... . -.-B-   .-Y-- ..- -.. .A",
			expected: "error trying to process morse code to text: invalid input(.... . -.-B-   .-Y-- ..- -.. .A). Only Morse code following the standard provided table should be entered",
		},
		{
			scenario: "Case 2",
			input:    "Ç.... . -.--  4 .--- F ..- G -.. . A",
			expected: "error trying to process morse code to text: invalid input(Ç.... . -.--  4 .--- F ..- G -.. . A). Only Morse code following the standard provided table should be entered",
		},
		{
			scenario: "Case 3",
			input:    " Ç .... . -.--  Z .--- F ..-  G  -.. . A ",
			expected: "error trying to process morse code to text: invalid input( Ç .... . -.--  Z .--- F ..-  G  -.. . A ). Only Morse code following the standard provided table should be entered",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			handler := NewMorseCodeHandler(dto.NewMorseCodeInputDTO(test.input))

			outputDTO, err := handler.Process()
			if err != nil && err.Error() != test.expected {
				t.Errorf("input: %q\nexpected: %q", test.input, test.expected)
			}

			if len(outputDTO.GetOutput()) > 0 {
				t.Errorf("input: %q\nexpected: %q\ngot: %q", test.input, "", outputDTO.GetOutput())
			}
		})
	}
}

func TestItShouldGenerateTheOutputWithoutErrorsEncapsulatedInTheDTO(t *testing.T) {
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
			handler := NewMorseCodeHandler(dto.NewMorseCodeInputDTO(test.input))

			outputDTO, _ := handler.Process()
			if outputDTO.GetOutput() != test.expected {
				t.Errorf("input: %q\nexpected: %q\ngot: %q", test.input, test.expected, outputDTO.GetOutput())
			}
		})
	}
}

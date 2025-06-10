package domain

import (
	"fmt"
	"strings"
)

var morseToChar = map[string]string{
	".-": "A", "-.": "N", "-----": "0",
	"-...": "B", "---": "O", ".----": "1",
	"-.-.": "C", ".--.": "P", "..---": "2",
	"-..": "D", "--.-": "Q", "...--": "3",
	".": "E", ".-.": "R", "....-": "4",
	"..-.": "F", "...": "S", ".....": "5",
	"--.": "G", "-": "T", "-....": "6",
	"....": "H", "..-": "U", "--...": "7",
	"..": "I", "...-": "V", "---..": "8",
	".---": "J", ".--": "W", "----.": "9",
	"-.-": "K", "-..-": "X",
	".-..": "L", "-.--": "Y",
	"--": "M", "--..": "Z",
}

type MorseCodeDomain struct {
	input  string
	output string
}

func NewMorseCodeDomain(input string) *MorseCodeDomain {
	return &MorseCodeDomain{
		input: input,
	}
}

func (mcd *MorseCodeDomain) Process() error {
	var builder strings.Builder

	phrases := strings.Split(mcd.input, "   ")
	for i, phrase := range phrases {
		characteres := strings.Split(phrase, " ")
		for _, key := range characteres {
			if value, ok := morseToChar[key]; ok {
				builder.WriteString(value)
			} else if !ok && len(key) > 0 {
				return fmt.Errorf("invalid input(%s). Only Morse code following the standard provided table should be entered", mcd.input)
			}
		}
		if i < len(phrases)-1 {
			builder.WriteString(" ")
		}
	}
	mcd.output = builder.String()
	return nil
}

func (mcd *MorseCodeDomain) GetOutput() string {
	return mcd.output
}

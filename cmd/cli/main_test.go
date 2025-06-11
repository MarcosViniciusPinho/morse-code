package main

import (
	"bytes"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestItSshouldReturnAnErrorWhenAttemptingToGenerateMessageFromInvalidMorseCode(t *testing.T) {

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
			origArgs := os.Args
			defer func() { os.Args = origArgs }()

			os.Args = []string{"morse-code", test.input}

			logs, _ := captureLogs(main)

			if !strings.Contains(logs, test.expected) {
				t.Fatalf("input: %q\nexpected field in logs: %q\nlogs: %s", test.input, test.expected, logs)
			}
		})
	}
}

func TestItShouldReturnParameterCountErrorWwhenAttemptingToGenerateMessageFromMorseCode(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
	}{
		{
			scenario: "Case 1",
			input:    ".... . -.--   .--- ..- -.. .",
		},
		{
			scenario: "Case 2",
			input:    ".... . .-.. .-.. ---   .-- --- .-. .-.. -..",
		},
		{
			scenario: "Case 3",
			input:    ".... . .-.. .-.. ---   .-- --- .-. .-.. -..   .... . -.--   .--- ..- -.. .",
		},
		{
			scenario: "Case 4",
			input:    " .... . -.--   .--- ..- -.. . ",
		},
		{
			scenario: "Case 5",
			input:    "  ....  .  -.--   .---  ..-  -..  .  ",
		},
		{
			scenario: "Case 6",
			input:    "  ....  .  -.--",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			origArgs := os.Args
			defer func() { os.Args = origArgs }()

			os.Args = []string{"morse-code", test.input, "YO"}

			logs, _ := captureLogs(main)

			expected := "invalid command or number of parameters, it should be used as follows:"

			if !strings.Contains(logs, expected) {
				t.Fatalf("input: %q\nexpected field in logs: %q\nlogs: %s", test.input, expected, logs)
			}
		})
	}
}

func TestItShouldReturnLogWithTheMorseCodeMessage(t *testing.T) {
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
			origArgs := os.Args
			defer func() { os.Args = origArgs }()

			os.Args = []string{"morse-code", test.input}

			logs, _ := captureLogs(main)

			if !strings.Contains(logs, test.expected) {
				t.Fatalf("input: %q\nexpected field in logs: %q\nlogs: %s", test.input, test.expected, logs)
			}
		})
	}
}

func captureLogs(f func()) (string, error) {
	orig := slog.Default()
	var buf bytes.Buffer
	handler := slog.NewTextHandler(&buf, nil)
	slog.SetDefault(slog.New(handler))
	f()
	slog.SetDefault(orig)
	return buf.String(), nil
}

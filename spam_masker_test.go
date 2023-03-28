package main

import "testing"

type spam struct {
	message  string
	expected string
}

var spamData = []spam{
	{
		message:  "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
		expected: "Here's my spammy page: http://******************* see you.",
	},
	{
		message:  "No spam here",
		expected: "No spam here",
	},
	{
		message:  "http://hehefouls.netHAHAHA",
		expected: "http://*******************",
	},
	{
		message:  "http://",
		expected: "http://",
	},
	{
		message:  "http://1",
		expected: "http://*",
	},
	{
		message:  "test",
		expected: "test",
	},
}

func TestSpamMasker(t *testing.T) {
	for _, test := range spamData {
		if output := spamMaker(test.message); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

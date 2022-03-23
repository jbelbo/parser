package parser

import "testing"

type isValidTestCase struct {
	text     string
	expected bool
}

var isValidTestCases = []isValidTestCase{
	{"[]", true},
	{"{}", true},
	{"()", true},
	{"({[]})", true},
	{"({[]{}()})", true},
	{"(aaaaa)", true},
	{"", true},
	{"a()", true},
	{"([)", false},
	{"([})", false},
	{"(}])", false},
}

func TestIsValid(t *testing.T) {
	for _, test := range isValidTestCases {
		if output := IsValid(test.text); output != test.expected {
			t.Errorf("Output %t for text %s not equal to expected %t", output, test.text, test.expected)
		}
	}
}

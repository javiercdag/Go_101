package strings

import (
	"unicode/utf8"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func multilineString() string {
	return "some\nmultiline\nstring"
}

func stringLen(s string) int {
	return utf8.RuneCountInString(s)
}

func trimFirstChar(s string) string {
	if stringLen(s) <= 0 {
		return s
	}

	_, size := utf8.DecodeRuneInString(s)

	return s[size:]
}

func trimLastChar(s string) string {
	if stringLen(s) <= 0 {
		return s
	}

	_, size := utf8.DecodeLastRuneInString(s)

	return s[:len(s)-size]
}

func swapFirstChar(s string) string {
	if stringLen(s) == 0 {
		return s
	}

	// "A" + string is optimized by the runtime
	return "A" + trimFirstChar(s)
}

func swapLastChar(s string) string {
	if len(s) == 0 {
		return s
	}

	return trimLastChar(s) + "A"
}

func prependChar(s string) string {
	return "A" + s
}

func appendChar(s string) string {
	return s + "A"
}

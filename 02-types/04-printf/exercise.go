package printer

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func printBool(b bool) string {
	// %t true false
	return fmt.Sprintf("variable of type boolean and value %t", b)
}

func printInt(i int) string {
	// %d decimal (base 10 integer)
	return fmt.Sprintf("variable of type integer and value %d", i)
}

func printHex(i int) string {
	// %x number to hexadecimal
	return fmt.Sprintf("variable of type integer in hexadecimal form and value %x", i)
}

func printFloat(f float64) string {
	// %f float"
	return fmt.Sprintf("variable of type float and value %.2f", f)
}

func printString(s string) string {
	// %s strings
	return fmt.Sprintf("variable of type string and value \"%s\"", s)
}

func concatStrings(a, b string) string {
	return a + b
}

func printConcatStrings(a, b string) string {
	return printString(concatStrings(a, b))
}

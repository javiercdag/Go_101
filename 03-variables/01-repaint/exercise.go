package repaint

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

var colorMap = map[string]string{
	"red":        "green",
	"green":      "red",
	"blue":       "orange",
	"orange":     "blue",
	"yellow":     "purple",
	"purple":     "yellow",
	"amber":      "violet",
	"violet":     "amber",
	"teal":       "vermilion",
	"vermilion":  "teal",
	"magenta":    "chartreuse",
	"chartreuse": "magenta",
}

// repaintColor returns the complementary of the color received as argument or an error for an unknown error.
func repaintColor(color string) (string, error) {
	val, ok := colorMap[color]

	if !ok {
		return "", fmt.Errorf("unknown color: %s", color)
	}

	return val, nil
}

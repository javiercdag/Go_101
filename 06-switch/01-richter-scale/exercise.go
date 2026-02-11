package richterscale

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// describeEarthquake returns the "description" of a given magnitude value on the Richter scale.
func describeEarthquake(magnitude float32) string {
	scale := ""

	switch {
	case magnitude < 2.0:
		scale = "micro"
	case magnitude < 3.0:
		scale = "very minor"
	case magnitude < 4.0:
		scale = "minor"
	case magnitude < 5.0:
		scale = "light"
	case magnitude < 6.0:
		scale = "moderate"
	case magnitude < 7.0:
		scale = "strong"
	case magnitude < 8.0:
		scale = "major"
	case magnitude < 10.0:
		scale = "great"
	case magnitude >= 10.0:
		scale = "massive"
	}

	return scale
}

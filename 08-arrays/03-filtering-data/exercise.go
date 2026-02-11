package filteringdata

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// filterData filters a slice based in an index slice.
func filterData(keys []string, indices []int) [10]string {
	resultingArray := [10]string{}
	counter := 0

	if len(keys) != len(indices) {
		return resultingArray
	}

	for i := 0; i < len(keys); i++ {
		if counter > 9 {
			return resultingArray
		}

		if indices[i]%2 != 0 {
			resultingArray[counter] = keys[i]
			counter++
		}
	}

	return resultingArray
}

package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt, isBroadcasted, isSecure bool) byte {
	var octet byte = (1 << 6) | (1 << 2)

	if isFirstAttempt {
		octet |= 1 << 4
	}

	if isBroadcasted {
		octet |= 1 << 1
	}

	if isSecure {
		octet |= 1 << 0
	}

	return octet
}

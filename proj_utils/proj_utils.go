package proj_utils

import (
	"bytes"
	"strconv"
)

func ArrayToString(A []int, delim string) string { //delim->delimiter

	var buffer bytes.Buffer
	for i := 0; i < len(A); i++ {
		buffer.WriteString(strconv.Itoa(A[i]))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

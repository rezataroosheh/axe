package extension

import (
	"fmt"
)

//Ternary method represents Ternary operator
func Ternary(condition bool, trueStatement, falseStatement interface{}) interface{} {
	if condition {
		return trueStatement
	}
	return falseStatement
}

//NewNilReferenceError return nil reference error
func NewNilReferenceError(parameterName string) error {
	return fmt.Errorf("Parameter name `%v` cannot be null", parameterName)
}

//Xor represents a function that xor part2 to part1
func Xor(part1 []byte, part2 []byte) {
	for i, b := range part2 {
		part1[i] ^= b
	}
}

// func LeftPad(data []byte, padding byte, totalLength int) []byte {
// 	if len(data) >= totalLength {
// 		return data
// 	}
// 	return append(bytes.Repeat([]byte{padding}, totalLength-len(data)), data...)
// }
// func RightPad(data []byte, padding byte, totalLength int) []byte {
// 	if len(data) >= totalLength {
// 		return data
// 	}
// 	return append(data, bytes.Repeat([]byte{padding}, totalLength-len(data))...)
// }

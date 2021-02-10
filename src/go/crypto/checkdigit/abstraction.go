package checkdigit

import (
	"errors"
	"strconv"
)

//CheckDigitAlgorithm represents an interface for check digit algorithms.
type CheckDigitAlgorithm interface {
	Compute(data []uint8) (uint8, error)
	GetName() string
}

//ComputeCheckDigit return Luhn (Modulus 10) algorithm check digit.
func ComputeCheckDigit(algorithm CheckDigitAlgorithm, str string) (uint8, error) {
	data := make([]uint8, len(str))
	for i, ch := range str {
		number, err := strconv.ParseInt(string(ch), 10, 16)
		if err != nil {
			return 255, err
		}
		data[i] = uint8(number)
	}
	return algorithm.Compute(data)
}

//VerifyCheckDigit validate the modulus 10 check digit.
func VerifyCheckDigit(algorithm CheckDigitAlgorithm, input string) (bool, uint8, error) {
	if len(input) < 2 {
		return false, 255, errors.New("Input is invalid")

	}
	lastIndex := len(input) - 1
	checkDigit, err := ComputeCheckDigit(algorithm, input[:lastIndex])
	if err != nil {
		return false, 255, err
	}
	//fmt.Printf("data is %v and check digit is %v. computed check digit is %v\n", valueWithCheckDigit, valueWithCheckDigit[:lastIndex], checkDigit)
	return strconv.Itoa(int(checkDigit)) == input[lastIndex:], checkDigit, nil
}

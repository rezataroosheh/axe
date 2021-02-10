package pinblock

import (
	"errors"
	"strings"
)

//OEM1PinBlockAlgorithm represnts The OEM-1 PIN block format is equivalent to the PIN block formats that Diebold, Docutel, and NCR define.
// The OEM-1 PIN block format supports a PIN from 4 to 12 digits in length.
// A PIN that is longer than 12 digits is truncated on the right.
type OEM1PinBlockAlgorithm struct {
	Padding PinBlockPadding
}

//NewFormat1 returns an instance of OEM1PinBlockAlgorithm
func NewOEM1(padding PinBlockPadding) *OEM1PinBlockAlgorithm {
	return &OEM1PinBlockAlgorithm{
		Padding: padding,
	}
}

//GetName returns OEM 1
func (algorithm OEM1PinBlockAlgorithm) GetName() string {
	return "OEM 1"
}

func (algorithm OEM1PinBlockAlgorithm) Encode(pin string) (string, error) {
	if _, err := validatePin(pin, 4, 12); err != nil {
		return "", err
	}

	padding := rune(algorithm.Padding)

	if strings.ContainsRune(pin, padding) {
		return "", errors.New("The pad value must be different from any PIN digit")
	}
	pin += strings.Repeat(string(padding), 16-len(pin))

	return pin, nil
}

func (algorithm OEM1PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	if len(pinblock) != 16 {
		return "", errors.New("Invalid length of pin block, length of Pin block must be 16")
	}
	return strings.TrimRight(pinblock, pinblock[len(pinblock)-1:]), nil
}

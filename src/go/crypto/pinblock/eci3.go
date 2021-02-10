package pinblock

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ECI3PinBlockAlgorithm represnts the ECI-3 PIN block format supports a PIN from 4 to 6 digits in length.
// A PIN that is longer than 6 digits is truncated on the right.
type ECI3PinBlockAlgorithm struct {
	GeneratePadding func(length int) []PinBlockPadding
}

//NewECI2 returns an instance of ECI3PinBlockAlgorithm
func NewECI3() *ECI3PinBlockAlgorithm {
	return &ECI3PinBlockAlgorithm{
		GeneratePadding: dedfaultPaddingGenerator,
	}
}

//GetName returns ECI-3
func (algorithm ECI3PinBlockAlgorithm) GetName() string {
	return "ECI-3"
}

func (algorithm ECI3PinBlockAlgorithm) Encode(pin string) (string, error) {
	if len(pin) < 4 || len(pin) > 6 {
		return "", errors.New("Invalid pin length, pin must be four digit")
	}
	pin = fmt.Sprintf("%v%s%s", len(pin), pin, strings.Repeat("0", 6-len(pin)))

	for _, pad := range algorithm.GeneratePadding(9) {
		pin += string(rune(pad))
	}

	return pin, nil
}
func (algorithm ECI3PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	if len(pinblock) != 16 {
		return "", errors.New("Invalid length of pin block, length of Pin block must be 16")
	}
	len, err := strconv.ParseInt(pinblock[0:1], 10, 16)
	if err != nil {
		return "", err
	}
	return pinblock[1:len], nil
}

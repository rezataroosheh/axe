package pinblock

import (
	"errors"
	"fmt"
	"strconv"
)

// Iso1PinBlockAlgorithm represnts ISO-1 (Format 1), that is ISO-1 PIN block format is equivalent to an ECI-4 PIN block format.
// The ISO-1 PIN block format supports a PIN from 4 to 12 digits in length.
// A PIN that is longer than 12 digits is truncated on the right.
type Iso1PinBlockAlgorithm struct {
	GeneratePadding func(length int) []PinBlockPadding
}

// NewFormat1 returns an instance of Iso1PinBlockAlgorithm
func NewFormat1() *Iso1PinBlockAlgorithm {
	return &Iso1PinBlockAlgorithm{GeneratePadding: dedfaultPaddingGenerator}
}

// GetName returns Format 1 (ISO-1)
func (algorithm Iso1PinBlockAlgorithm) GetName() string {
	return "Format 1 (ISO-1)"
}

func (algorithm Iso1PinBlockAlgorithm) Encode(pin string) (string, error) {
	if _, err := validatePin(pin, 4, 14); err != nil {
		return "", err
	}
	pin = fmt.Sprintf("1%X%s", len(pin), pin)

	pin += convert(algorithm.GeneratePadding(16 - len(pin)))

	return pin, nil
}
func (algorithm Iso1PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	var (
		err error
	)
	if len(pinblock) != 16 {
		return "", errors.New("Invalid length of pin block, length of Pin block must be 16")
	}
	len, err := strconv.ParseInt(pinblock[1:2], 16, 16)
	if err != nil {
		return "", err
	}
	return pinblock[2 : 2+len], nil
}

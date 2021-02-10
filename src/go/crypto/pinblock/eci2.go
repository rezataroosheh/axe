package pinblock

import (
	"errors"
)

// ECI2PinBlockAlgorithm represnts the ECI-2 PIN block format supports a 4-digit PIN.
// A PIN that is longer than 4 digits is truncated on the right.
type ECI2PinBlockAlgorithm struct {
	GeneratePadding func(length int) []PinBlockPadding
}

//NewECI2 returns an instance of ECI2PinBlockAlgorithm
func NewECI2() *ECI2PinBlockAlgorithm {
	return &ECI2PinBlockAlgorithm{
		GeneratePadding: dedfaultPaddingGenerator,
	}
}

//GetName returns ECI-2
func (algorithm ECI2PinBlockAlgorithm) GetName() string {
	return "ECI-2"
}

func (algorithm ECI2PinBlockAlgorithm) Encode(pin string) (string, error) {
	if len(pin) != 4 {
		return "", errors.New("Invalid pin length, pin must be four digit")
	}
	for _, pad := range algorithm.GeneratePadding(12) {
		pin += string(rune(pad))
	}

	return pin, nil
}
func (algorithm ECI2PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	if len(pinblock) != 16 {
		return "", errors.New("Invalid length of pin block, length of Pin block must be 16")
	}
	return pinblock[:4], nil
}

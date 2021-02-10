package pinblock

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Iso2PinBlockAlgorithm represnts ISO 2 (Format 2), that is for local use with off-line systems only (e.g. smart cards).
type Iso2PinBlockAlgorithm struct {
	Padding PinBlockPadding
}

// NewFormat1 returns an instance of Iso2PinBlockAlgorithm
func NewFormat2(padding PinBlockPadding) *Iso2PinBlockAlgorithm {
	return &Iso2PinBlockAlgorithm{
		Padding: padding,
	}
}

// GetName returns Format 2 (ISO-2)
func (algorithm Iso2PinBlockAlgorithm) GetName() string {
	return "Format 2 (ISO-2)"
}

func (algorithm Iso2PinBlockAlgorithm) Encode(pin string) (string, error) {
	if _, err := validatePin(pin, 4, 14); err != nil {
		return "", err
	}
	pin = fmt.Sprintf("2%X%s%s", len(pin), pin, strings.Repeat(string(algorithm.Padding), 16-len(pin)))

	return pin, nil
}
func (algorithm Iso2PinBlockAlgorithm) Decode(pinblock string) (string, error) {
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

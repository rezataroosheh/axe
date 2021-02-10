package pinblock

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rezataroosheh/axe/internal/extension"
)

// Iso0PinBlockAlgorithm represents ISO-0 (Format 0), thats is the first and most common PIN block encoding format based on ISO 9564 – an international
// standard for personal identification number (PIN) management and security in retail banking.
type Iso0PinBlockAlgorithm struct {
	Pan     string
	Padding PinBlockPadding
}

func NewFormat0(pan string, padding PinBlockPadding) *Iso0PinBlockAlgorithm {
	return &Iso0PinBlockAlgorithm{
		Pan:     pan,
		Padding: padding,
	}
}

//GetName returns Format 0 (ISO-0)
func (algorithm Iso0PinBlockAlgorithm) GetName() string {
	return "Format 0 (ISO-0)"
}

func (algorithm Iso0PinBlockAlgorithm) Encode(pin string) (string, error) {
	var (
		panPart, pinPart []byte
		err              error
		pan              = algorithm.Pan
		padding          = algorithm.Padding
		lastIndex        = len(pan) - 1
	)
	if valid, err := validatePan(pan); !valid {
		return "", err
	}
	if _, err = validatePin(pin, 4, 14); err != nil {
		return "", err
	}

	pan = fmt.Sprintf("0000%s", pan[lastIndex-12:lastIndex])
	pin = fmt.Sprintf("0%x%s", len(pin), pin)
	pin += strings.Repeat(string(padding), 16-len(pin))

	if panPart, err = hex.DecodeString(pan); err != nil {
		return "", err
	}
	if pinPart, err = hex.DecodeString(pin); err != nil {
		return "", err
	}
	extension.Xor(panPart, pinPart)
	return hex.EncodeToString(panPart), nil
}
func (algorithm Iso0PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	var (
		panPart, pinblockPart []byte
		err                   error
		pan                   = algorithm.Pan
		lastIndex             = len(pan) - 1
	)
	if valid, err := validatePan(pan); !valid {
		return "", err
	}
	if len(pinblock) != 16 {
		return "", errors.New("Invalid length of pin block, length of Pin block must be 16")
	}
	pan = fmt.Sprintf("0000%s", pan[lastIndex-12:lastIndex])

	if panPart, err = hex.DecodeString(pan); err != nil {
		return "", err
	}
	if pinblockPart, err = hex.DecodeString(pinblock); err != nil {
		return "", err
	}
	extension.Xor(panPart, pinblockPart)

	result := hex.EncodeToString(panPart)
	len, err := strconv.ParseInt(result[1:2], 16, 16)
	if err != nil {
		return "", err
	}
	return result[2 : 2+len], nil
}

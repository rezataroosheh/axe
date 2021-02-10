package pinblock

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/rezataroosheh/axe/internal/extension"
)

// Iso3PinBlockAlgorithm represnts ISO-3 (Format 3), that is Format 3 is the same as format 0, except that the “fill”
// digits are random values from 10 to 15, and the first nibble (which identifies the block format) has the value 3.
type Iso3PinBlockAlgorithm struct {
	GeneratePadding func(length int) []PinBlockPadding
	Pan             string
}

//NewFormat3 returns an instance of Iso3PinBlockAlgorithm
func NewFormat3(pan string) *Iso3PinBlockAlgorithm {
	return &Iso3PinBlockAlgorithm{
		GeneratePadding: dedfaultFormat3GeneratePadding,
		Pan:             pan,
	}
}

func dedfaultFormat3GeneratePadding(length int) []PinBlockPadding {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	result := make([]PinBlockPadding, length)
	for i := 0; i < length; i++ {
		value := rnd.Intn(5) + 10
		result[i] = convertToPinBlock(value)
	}
	return result
}

//GetName returns Format 3 (ISO-3)
func (algorithm Iso3PinBlockAlgorithm) GetName() string {
	return "Format 3 (ISO-3)"
}

func (algorithm Iso3PinBlockAlgorithm) Encode(pin string) (string, error) {
	var (
		panPart, pinPart []byte
		err              error
		pan              = algorithm.Pan
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

	for _, pad := range algorithm.GeneratePadding(16 - len(pin)) {
		if pad < A {
			return "", errors.New("The padding must be between 10 and 15")
		}
		pin += string(rune(pad))
	}

	if panPart, err = hex.DecodeString(pan); err != nil {
		return "", err
	}
	if pinPart, err = hex.DecodeString(pin); err != nil {
		return "", err
	}
	extension.Xor(panPart, pinPart)
	return hex.EncodeToString(panPart), nil
}

func (algorithm Iso3PinBlockAlgorithm) Decode(pinblock string) (string, error) {
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

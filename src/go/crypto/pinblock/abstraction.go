package pinblock

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PinBlockPadding rune

const (
	Zero  PinBlockPadding = '0'
	One   PinBlockPadding = '1'
	Two   PinBlockPadding = '2'
	Three PinBlockPadding = '3'
	Four  PinBlockPadding = '4'
	Five  PinBlockPadding = '5'
	Six   PinBlockPadding = '6'
	Seven PinBlockPadding = '7'
	Eight PinBlockPadding = '8'
	Nine  PinBlockPadding = '9'
	A     PinBlockPadding = 'A'
	B     PinBlockPadding = 'B'
	C     PinBlockPadding = 'C'
	D     PinBlockPadding = 'D'
	E     PinBlockPadding = 'E'
	F     PinBlockPadding = 'F'
)

var (
	paddingMap = map[int]PinBlockPadding{
		0:  Zero,
		1:  One,
		2:  Two,
		3:  Three,
		4:  Four,
		5:  Five,
		6:  Six,
		7:  Seven,
		8:  Eight,
		9:  Nine,
		10: A,
		11: B,
		12: C,
		13: D,
		14: E,
		15: F,
	}
)

//PinBlockAlgorithm represents an interface for computing pin blocks algorithms.
type PinBlockAlgorithm interface {
	Encode(pin string) (string, error)
	Decode(pinblock string) (string, error)
	GetName() string
}

func convertToPinBlock(value int) PinBlockPadding {
	return paddingMap[value]
}

func validatePan(pan string) (bool, error) {
	if len(pan) < 12 || len(pan) > 19 {
		return false, errors.New("Invalid pan length, the length of PAN must be between 12 and 19")
	}
	return true, nil
}
func validatePin(pin string, minLen, maxLen int) (bool, error) {
	if len(pin) < minLen || len(pin) > maxLen {
		return false, fmt.Errorf("Invalid pin length, the length of PIN must be between %d and %d", minLen, maxLen)
	}
	return true, nil
}

func dedfaultPaddingGenerator(length int) []PinBlockPadding {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	result := make([]PinBlockPadding, length)
	for i := 0; i < length; i++ {
		value := rnd.Intn(16)
		result[i] = convertToPinBlock(value)
	}
	return result
}
func convert(paddings []PinBlockPadding) string {
	var str string
	for _, pad := range paddings {
		str += string(rune(pad))
	}
	return str
}

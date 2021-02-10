package checkdigit

import "errors"

//Modulus10Algorithm represents Luhn algorithm known as Modulus 10 algorithm.
type Modulus10Algorithm struct {
}

//GetName return modulus10 as name of algorithm.
func (algorithm Modulus10Algorithm) GetName() string {
	return "modulus10"
}

//ComputeCheckDigit return Luhn (Modulus 10) algorithm check digit.
func (algorithm Modulus10Algorithm) Compute(data []uint8) (uint8, error) {
	sum := 0
	parity := len(data) % 2
	for i, number := range data {
		if number > 9 {
			return 255, errors.New("The numbers must be less than 10")
		}
		value := int(number) * (1 + (i % 2) ^ parity)

		if value > 9 {
			value -= 9
		}
		sum += value
	}
	sum %= 10
	if sum == 0 {
		return 0, nil
	}

	return uint8(10 - sum), nil
}

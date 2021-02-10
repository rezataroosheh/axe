package checkdigit

import "errors"

//VerhoeffAlgorithm represents Luhn algorithm known as Verhoeff algorithm.
type VerhoeffAlgorithm struct {
}

var (
	multiplicationTable = [][]uint8{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
		{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
		{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
		{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
		{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
		{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
		{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
		{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}

	permutationTable = [][]uint8{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
		{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
		{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
		{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
		{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
		{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
		{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
	}
	inverseTable = []uint8{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}
)

//GetName return verhoeff as name of algorithm.
func (algorithm VerhoeffAlgorithm) GetName() string {
	return "verhoeff"
}

//ComputeCheckDigit return Luhn (Modulus 10) algorithm check digit.
func (algorithm VerhoeffAlgorithm) Compute(data []uint8) (uint8, error) {
	var (
		index uint8 = 0
		len   int   = len(data)
	)
	for i := 0; i < len; i++ {
		if data[i] > 9 {
			return 255, errors.New("The numbers must be less than 10")
		}
		firstIndex := (i + 1) % 8
		secondIndex := data[len-i-1]
		permutation := permutationTable[firstIndex][secondIndex]
		index = multiplicationTable[index][permutation]
	}
	return inverseTable[index], nil
}

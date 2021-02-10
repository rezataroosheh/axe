package test

import (
	"testing"

	"github.com/rezataroosheh/axe/crypto/checkdigit"
)

type CheckDigitAlgorithmTestCase struct {
	Algorithm                 checkdigit.CheckDigitAlgorithm
	Name, AssertMessage, Data string
}

var (
	checkDigitAlgorithmTestCases = []CheckDigitAlgorithmTestCase{
		{
			Name:          "Modulus 10 test case.",
			Algorithm:     checkdigit.Modulus10Algorithm{},
			Data:          "79927398713",
			AssertMessage: "Expected check digit is %v but computed is %v",
		},
		{
			Name:          "Verhoeff test case.",
			Algorithm:     checkdigit.VerhoeffAlgorithm{},
			Data:          "12345678902",
			AssertMessage: "Expected check digit is %v but computed is %v",
		},
	}
)

func TestAllCheckDigitAlgorithm(t *testing.T) {
	t.Parallel()

	for _, testCase := range checkDigitAlgorithmTestCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result, checkDigit, err := checkdigit.VerifyCheckDigit(testCase.Algorithm, testCase.Data)
			if err != nil {
				t.Error(err)
				return
			}
			if !result {
				t.Errorf(testCase.AssertMessage, testCase.Data, checkDigit)
				return
			}
		})
	}
}

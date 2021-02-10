package test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/rezataroosheh/axe/crypto/ansimac"
)

type AnsiMacAlgorithmTestCase struct {
	Algorithm           ansimac.AnsiMacAlgorithm
	Data, Expected      []byte
	HappyScenario       bool
	Name, AssertMessage string
}

var (
	key1      = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	key2      = []byte{0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	key3      = []byte{0xC1, 0xD0, 0xF8, 0xFB, 0x49, 0x58, 0x67, 0x0D}
	noPadData = []byte{0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x69, 0x6D, 0x65, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x61, 0x6C, 0x6C, 0x20, 0xAB, 0xCD}
	padData   = []byte{0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x69, 0x6D, 0x65, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x61, 0x6C, 0x6C, 0x20}

	expectedAnsi99NoPadData = []byte{0x36, 0x61, 0x1D, 0xBB, 0x2D, 0x0A, 0xC1, 0xE6}
	expectedAnsi99PadData   = []byte{0x70, 0xA3, 0x06, 0x40, 0xCC, 0x76, 0xDD, 0x8B}

	expectedAnsi919NoPadData = []byte{0x1C, 0x05, 0x08, 0x79, 0xD9, 0x58, 0x16, 0xB8}
	expectedAnsi919PadData   = []byte{0xA1, 0xC7, 0x2E, 0x74, 0xEA, 0x3F, 0xA9, 0xB6}

	expectedRetailPadData   = []byte{0x31, 0x97, 0x8A, 0xA2, 0xFB, 0x5A, 0xC7, 0x4E}
	expectedRetailNoPadData = []byte{0xCD, 0x1D, 0x84, 0xDD, 0x37, 0xBC, 0xEF, 0xFD}

	nilEror           = "Bad result: error must be filed when one of parameters nil"
	badMacComputation = "Bad result: %v. The result must be %v"

	ansi99    = ansimac.Ansi99MacAlgorithm{Key: key1}
	ansi919   = ansimac.Ansi919MacAlgorithm{Key1: key1, Key2: key2}
	retailmac = ansimac.RetailMacAlgorithm{Key1: key1, Key2: key2, Key3: key3}

	ansiMacAlgorithmTestCases = []AnsiMacAlgorithmTestCase{
		{
			Name:          "Ansi 9.9 MAC test case for not pad data.",
			Algorithm:     ansi99,
			Data:          noPadData,
			Expected:      expectedAnsi99NoPadData,
			HappyScenario: true,
			AssertMessage: badMacComputation,
		},
		{
			Name:          "Ansi 9.9 MAC test case for pad data.",
			Algorithm:     ansi99,
			Data:          padData,
			Expected:      expectedAnsi99PadData,
			HappyScenario: true,
			AssertMessage: badMacComputation,
		},
		{
			Name:          "Ansi 9.9 MAC test case for nil data",
			Algorithm:     ansi99,
			Data:          nil,
			Expected:      expectedAnsi99NoPadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
		{
			Name:          "Ansi 9.9 MAC test case for nil key",
			Algorithm:     ansimac.Ansi99MacAlgorithm{Key: nil},
			Data:          padData,
			Expected:      expectedAnsi919PadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},

		{
			Name:          "Ansi 9.19 MAC test case for pad data.",
			Algorithm:     ansi919,
			Data:          padData,
			Expected:      expectedAnsi919PadData,
			HappyScenario: true,
			AssertMessage: badMacComputation,
		},
		{
			Name:          "Ansi 9.19 MAC test case for not pad data.",
			Algorithm:     ansi919,
			Data:          noPadData,
			Expected:      expectedAnsi919NoPadData,
			HappyScenario: true,
			AssertMessage: badMacComputation,
		},
		{
			Name:          "Ansi 9.19 MAC test case for nil Key1",
			Algorithm:     ansimac.Ansi919MacAlgorithm{Key1: nil, Key2: key2},
			Data:          padData,
			Expected:      expectedAnsi919PadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
		{
			Name:          "Ansi 9.19 MAC test case for nil Key2",
			Algorithm:     ansimac.Ansi919MacAlgorithm{Key1: key1, Key2: nil},
			Data:          padData,
			Expected:      expectedAnsi919PadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
		{
			Name:          "Ansi 9.19 MAC test case for nil data.",
			Algorithm:     ansi919,
			Data:          nil,
			Expected:      expectedAnsi919NoPadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},

		{
			Name:          "Retail MAC test case for pad data.",
			Algorithm:     retailmac,
			Data:          padData,
			Expected:      expectedRetailPadData,
			HappyScenario: true,
			AssertMessage: badMacComputation,
		},
		{
			Name:          "Retail MAC test case for not pad data.",
			Algorithm:     retailmac,
			Data:          noPadData,
			Expected:      expectedRetailNoPadData,
			HappyScenario: true,
			AssertMessage: badMacComputation,
		},
		{
			Name:          "Retail MAC test case for nil Key1",
			Algorithm:     ansimac.RetailMacAlgorithm{Key1: nil, Key2: key2, Key3: key3},
			Data:          padData,
			Expected:      expectedRetailPadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
		{
			Name:          "Retail MAC test case for nil Key2",
			Algorithm:     ansimac.RetailMacAlgorithm{Key1: key1, Key2: nil, Key3: key3},
			Data:          padData,
			Expected:      expectedRetailPadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
		{
			Name:          "Retail MAC test case for nil Key3",
			Algorithm:     ansimac.RetailMacAlgorithm{Key1: key1, Key2: key2, Key3: nil},
			Data:          padData,
			Expected:      expectedRetailPadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
		{
			Name:          "Retail MAC test case for nil data.",
			Algorithm:     retailmac,
			Data:          nil,
			Expected:      expectedRetailPadData,
			HappyScenario: false,
			AssertMessage: nilEror,
		},
	}
)

func TestAllAnsiMacAlgorithm(t *testing.T) {
	t.Parallel()

	for _, testCase := range ansiMacAlgorithmTestCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			result, err := testCase.Algorithm.Compute(testCase.Data)

			if err != nil {
				if testCase.HappyScenario {
					t.Error(err)
				}
				return
			}
			if !testCase.HappyScenario {
				t.Errorf(testCase.AssertMessage)
				return
			}
			if !bytes.Equal(testCase.Expected, result) {
				t.Errorf(testCase.AssertMessage, hex.EncodeToString(result), hex.EncodeToString(testCase.Expected))
				return
			}
		})
	}
}

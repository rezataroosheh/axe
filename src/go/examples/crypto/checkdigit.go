package crypto

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rezataroosheh/axe/crypto/checkdigit"
)

//CheckDigitExample is an example
type CheckDigitExample struct {
}

//GetName return name of example
func (example *CheckDigitExample) GetName() string {
	return "check-digit"
}

//Execute the exmple
func (example *CheckDigitExample) Execute(commands []string) bool {
	if commands == nil {
		return false
	}
	var (
		commandFlag   = flag.NewFlagSet(example.GetName(), flag.ExitOnError)
		modeFlag      = commandFlag.String("m", "verify", "Mode {verify|compute}")
		algorithmFlag = commandFlag.String("a", "modulus10", "Algorithm {modulus10|verhoeff}")
		inputFlag     = commandFlag.String("i", "", "Input value")
		showHelpFlag  = commandFlag.Bool("h", false, "Show help message")
	)
	if err := commandFlag.Parse(commands); err != nil {
		return false
	}

	if *showHelpFlag {
		commandFlag.Usage()
		return false
	}
	algorithmChoices := map[string]checkdigit.CheckDigitAlgorithm{
		"modulus10": checkdigit.Modulus10Algorithm{},
		"verhoeff":  checkdigit.VerhoeffAlgorithm{},
	}

	modeChoices := map[string]func(algorithm checkdigit.CheckDigitAlgorithm, input string) bool{
		"verify":  verify,
		"compute": compute,
	}

	algorithmChoice, validAlgorithmChoice := algorithmChoices[strings.ToLower(*algorithmFlag)]
	action, validModeChoice := modeChoices[strings.ToLower(*modeFlag)]

	if !validModeChoice || !validAlgorithmChoice {
		commandFlag.Usage()
		return false
	}
	return action(algorithmChoice, *inputFlag)
}

func verify(algorithm checkdigit.CheckDigitAlgorithm, input string) bool {
	result, checkDigit, err := checkdigit.VerifyCheckDigit(algorithm, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check digit computation error: %v\n", err)
		return false
	}
	if !result {
		fmt.Printf("Check digit for %v is %v but expected is %v.", input, checkDigit, input[len(input)-1:])
		return false
	}
	fmt.Printf("Check digit validated for %v.", input)
	return true
}
func compute(algorithm checkdigit.CheckDigitAlgorithm, input string) bool {
	result, err := checkdigit.ComputeCheckDigit(algorithm, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check digit computation error: %v\n", err)
		return false
	}
	fmt.Printf("Check digit for %v with algorithm %v is %v.", input, algorithm.GetName(), result)
	return true
}

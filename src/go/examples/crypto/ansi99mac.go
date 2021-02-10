package crypto

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/rezataroosheh/axe/crypto/ansimac"
)

//Ansi99MacExample is an example
type Ansi99MacExample struct {
}

//GetName return name of example
func (example *Ansi99MacExample) GetName() string {
	return "ansi-9.9"
}

//Execute the exmple
func (example *Ansi99MacExample) Execute(commands []string) bool {
	if commands == nil {
		return false
	}
	var (
		commandFlag  = flag.NewFlagSet(example.GetName(), flag.ExitOnError)
		keyFlag      = commandFlag.String("k", "", "Algorithm key")
		inputFlag    = commandFlag.String("i", "", "Input value")
		showHelpFlag = commandFlag.Bool("h", false, "Show help message")

		err            error
		key, mac, data []byte
	)
	if err := commandFlag.Parse(commands); err != nil {
		return false
	}

	if *showHelpFlag {
		commandFlag.Usage()
		return false
	}

	if key, err = hex.DecodeString(*keyFlag); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Key: %v\n", err)
		return false
	}
	ansi99 := ansimac.Ansi99MacAlgorithm{Key: key}

	if data, err = hex.DecodeString(*inputFlag); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Data: %v\n", err)
		return false
	}
	if mac, err = ansi99.Compute(data); err != nil {
		fmt.Fprintf(os.Stderr, "mac calculate: %v\n", err)
		return false
	}
	fmt.Printf("Mac computation result-> %v\n", hex.EncodeToString(mac))
	return true
}

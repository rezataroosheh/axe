package crypto

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/rezataroosheh/axe/crypto/ansimac"
)

//Ansi919MacExample is an example
type Ansi919MacExample struct {
}

//GetName return name of example
func (example *Ansi919MacExample) GetName() string {
	return "ansi-9.19"
}

//Execute the exmple
func (example *Ansi919MacExample) Execute(commands []string) bool {
	if commands == nil {
		return false
	}
	var (
		commandFlag  = flag.NewFlagSet(example.GetName(), flag.ExitOnError)
		key1Flag     = commandFlag.String("k1", "", "Algorithm first key")
		key2Flag     = commandFlag.String("k2", "", "Algorithm second key")
		inputFlag    = commandFlag.String("i", "", "Input value")
		showHelpFlag = commandFlag.Bool("h", false, "Show help message")

		err                   error
		key1, key2, mac, data []byte
	)
	if err := commandFlag.Parse(commands); err != nil {
		return false
	}

	if *showHelpFlag {
		commandFlag.Usage()
		return false
	}

	if key1, err = hex.DecodeString(*key1Flag); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Key: %v\n", err)
		return false
	}
	if key2, err = hex.DecodeString(*key2Flag); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Key: %v\n", err)
		return false
	}
	ansi919 := ansimac.Ansi919MacAlgorithm{Key1: key1, Key2: key2}

	if data, err = hex.DecodeString(*inputFlag); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Data: %v\n", err)
		return false
	}
	if mac, err = ansi919.Compute(data); err != nil {
		fmt.Fprintf(os.Stderr, "mac calculate: %v\n", err)
		return false
	}
	fmt.Printf("Mac computation result-> %v\n", hex.EncodeToString(mac))
	return true
}

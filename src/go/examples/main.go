package main

import (
	"github.com/rezataroosheh/axe/examples/crypto"
	"github.com/rezataroosheh/axe/internal/exampleutils"
)

func main() {
	exampleRepresenters := []exampleutils.ExampleRepresenter{
		&crypto.CheckDigitExample{},
		&crypto.Ansi99MacExample{},
		&crypto.Ansi919MacExample{},
	}
	exampleutils.RunExamples(exampleRepresenters)
}

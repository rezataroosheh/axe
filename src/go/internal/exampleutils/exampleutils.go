package exampleutils

import (
	"fmt"
	"os"
	"strings"
)

//ExampleRepresenter provide an interface to represents exmaples
type ExampleRepresenter interface {
	GetName() string
	Execute(commands []string) bool
}

//RunExamples executes examples pipe
func RunExamples(exampleRepresenters []ExampleRepresenter) {
	examplemetadatas := makeCommandsMetaDatas(exampleRepresenters)

	if len(os.Args) < 2 {
		showCommandsAndExit(examplemetadatas, 1)
		return
	}
	choice, valid := examplemetadatas[strings.ToLower(os.Args[1])]
	if !valid {
		showCommandsAndExit(examplemetadatas, 1)
		return
	}
	if !choice.Execute(os.Args[2:]) {
		os.Exit(1)
	}
}
func makeCommandsMetaDatas(exampleRepresenters []ExampleRepresenter) map[string]ExampleRepresenter {
	examplemetadatas := make(map[string]ExampleRepresenter)
	for _, example := range exampleRepresenters {
		examplemetadatas[example.GetName()] = example
	}
	return examplemetadatas
}

func showCommandsAndExit(items map[string]ExampleRepresenter, exitcode int) {
	fmt.Println("Valid commands are:")
	for key := range items {
		fmt.Printf("\t%v\n", key)
	}
	os.Exit(exitcode)
}

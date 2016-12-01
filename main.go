package main

import (
	"log"
	"os"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	args, err := docopt.Parse(usage, nil, true, "i3x "+version, false)
	if err != nil {
		log.Fatalln(err)
	}

	// log.Println("Arguments:", args)

	// Run action
	switch {
	case args["output"].(bool):
		if args["switch"].(bool) {
			outputSwitch()
		}
	}

	os.Exit(0)
}

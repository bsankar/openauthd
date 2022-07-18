package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	var CommandLine = flag.NewFlagSet(args[0], flag.ExitOnError)
	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	if len(args) != 2 {
		flag.Usage()
	}

	switch args[1] {
	case "init":
		OadInit()
	case "start":
	}

}

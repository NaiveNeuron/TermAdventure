package main

import (
	"./levels"
	"flag"
	"fmt"
	"os"
)

func main() {
	print_flag := flag.Bool("print", false, "print loaded levels and exit")
	pretty_print_flag := flag.Bool("no-pretty-print", false,
		"disable option to skip pretty printing")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("\n\nNo input file\n\n")
		fmt.Printf("usage: %s path\n\n", os.Args[0])
		os.Exit(0)
	}
	path := flag.Args()[0]
	challenge_name := levels.BasenameFromPath(path)

	challenge := levels.NewChallenge(challenge_name)

	challenge.LoadFromFile(path)

	if *print_flag {
		challenge.Print()
		os.Exit(0)
	}

	challenge.LoadCfg()

	if challenge.CheckCurrentLevel() {
		challenge.PrintCurrentLevel(*pretty_print_flag)
		challenge.IncreaseLevel()
	}

	challenge.PrintIdentifier()
}

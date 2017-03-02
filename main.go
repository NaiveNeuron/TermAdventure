package main

import (
	"./levels"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	key := []byte("example key 1234")

	print_flag := flag.Bool("print", false, "print loaded levels and exit")
	pretty_print_flag := flag.Bool("no-pretty-print", false,
		"disable option to skip pretty printing")
	encrypt_flag := flag.Bool("enc", false, "encrypt a given challenge")
	decrypt_flag := flag.Bool("dec", false, "decrypt a given challenge")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("\n\nNo input file\n\n")
		fmt.Printf("usage: %s path\n\n", os.Args[0])
		os.Exit(0)
	}
	path := flag.Args()[0]
	challenge_name := levels.BasenameFromPath(path)

	challenge := levels.NewChallenge(challenge_name)
	challenge_text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	if *encrypt_flag {
		encrypted_text, err := levels.Encrypt(key, string(challenge_text))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encrypted_text)
		os.Exit(0)
	}

	decrypted_text, err := levels.Decrypt(key, string(challenge_text))
	if err != nil {
		log.Fatal(err)
	}

	if *decrypt_flag {
		fmt.Println(decrypted_text)
		os.Exit(0)
	}

	challenge.LoadFromString(decrypted_text)

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

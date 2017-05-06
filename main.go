package main

import (
	"./levels"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var encryption_key string

func main() {
	key := []byte(encryption_key)

	print_flag := flag.Bool("print", false, "print loaded levels and exit")
	pretty_print_flag := flag.Bool("no-pretty-print", false,
		"disable option to skip pretty printing")
	detect_level_flag := flag.Bool("detect-level", false,
		"detect a level from a given hash and home directory")
	encrypt_flag := flag.Bool("enc", false, "encrypt a given challenge")
	decrypt_flag := flag.Bool("dec", false, "decrypt a given challenge")
	template_flag := flag.Bool("generate-from-template", false,
		"generate content of a .gta file from given template file and variables file")
	flag.BoolVar(template_flag, "g", false,
		"generate content of a .gta file from given template file and variables file")
	print_identifier_flag := flag.Bool("print-identifier", false,
		"print level identifier and exit")
	print_sleep_time := flag.Int("print-sleep-time", 50,
		"sets the amount of milliseconds that need to occur between two chars being printed")
	print_current_level_flag := flag.Bool("print-current-level", false,
		"print current level and exit")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("\n\nNo input file\n\n")
		if *template_flag {
			fmt.Printf("usage: %s template_file [variables_file]\n\n"+
				"variables_file=template_name.yaml by default\n", os.Args[0])
		} else {
			fmt.Printf("usage: %s gta_file>\n", os.Args[0])
		}
		os.Exit(1)
	}

	if *template_flag {
		templ, err := ioutil.ReadFile(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n\n", err)
			fmt.Printf("usage: %s template_file [variables_file]\n\n"+
				"variables_file=template_name.yaml by default\n", os.Args[0])
			os.Exit(1)
		}
		var yaml_name string
		if len(flag.Args()) >= 2 {
			yaml_name = flag.Args()[1]
		} else {
			yaml_name = levels.BasenameFromPath(flag.Args()[0]) + ".yaml"
		}
		yaml_data, yaml_err := ioutil.ReadFile(yaml_name)
		if yaml_err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", yaml_err)
			fmt.Printf("usage: %s template_file [variables_file]\n\n"+
				"variables_file=template_name.yaml by default\n", os.Args[0])
			os.Exit(1)
		}
		levels.Template(templ, yaml_data)
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

	decrypted_text := string(challenge_text)
	if strings.HasSuffix(path, ".enc") {
		decrypted_text, err = levels.Decrypt(key, string(challenge_text))
		if err != nil {
			log.Fatal(err)
		}
	}

	if *decrypt_flag {
		fmt.Println(decrypted_text)
		os.Exit(0)
	}

	challenge.LoadFromString(decrypted_text)

	if *detect_level_flag {
		if len(flag.Args()) < 3 {
			fmt.Printf("usage: %s path hash homedir\n", os.Args[0])
			os.Exit(1)
		}

		level := flag.Args()[1]
		homedir := flag.Args()[2]
		str, i := challenge.IDAndHomedirToLevel(level, homedir)
		if i != -1 {
			fmt.Println("Detected level:", str)
			os.Exit(0)
		} else {
			fmt.Println("Level undetected")
			os.Exit(1)
		}
	}

	if *print_flag {
		challenge.Print()
		os.Exit(0)
	}

	challenge.LoadCfg()

	if *print_current_level_flag {
		challenge.PrintCurrentLevel(*pretty_print_flag, *print_sleep_time)
		os.Exit(0)
	}

	if *print_identifier_flag {
		challenge.PrintIdentifier()
		os.Exit(0)
	}

	if challenge.CheckCurrentLevel() {
		challenge.GoToNextLevel()
	}

	// Print the current level definition either if we have just switched levels,
	// or if the user has explicitly requested that.
	print_again_exists, _ := levels.CmdOK("test -e $HOME/.gta_print_again")
	if *challenge.LastLevelPrinted != "yes" || print_again_exists {
		challenge.PrintCurrentLevel(*pretty_print_flag, *print_sleep_time)

		// Make sure that the level definition won't be printed again,
		// unless the user has done any action that suggests it should.
		challenge.SetConfigVal("last_level_printed", "yes")
		levels.CmdOK("rm -f $HOME/.gta_print_again")
	}
}

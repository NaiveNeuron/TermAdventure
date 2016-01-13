package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"os/exec"
)

func cmd_ok(cmd string) bool {
	_, err := exec.Command("sh", "-c", cmd).Output()
	return err == nil
}

func main() {
	mypath, _ := osext.Executable()
	exec.Command("env", "PROMPT_COMMAND="+mypath)

	if cmd_ok("[[ $(pwd) == \"$HOME\" ]]") {
		fmt.Println("At Home!")
	}
}

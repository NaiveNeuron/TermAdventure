package levels

import (
	"os/exec"
	"strings"
	"time"
)

func CmdOK(cmd string) bool {
	_, err := exec.Command("sh", "-c", cmd).Output()
	return err == nil
}

func print_line(text string) {
	for _, char := range text {
		print(string(char))
		time.Sleep(50 * time.Millisecond)
		if char == '.' || char == '!' || char == '?' {
			time.Sleep(500 * time.Millisecond)
		}
	}
	print("\n")
}

func PrintText(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		print_line(line)
	}
}

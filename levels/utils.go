package levels

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

func CmdOK(cmd string) (bool, string) {
	if cmd == "" {
		return true, ""
	}
	output, err := exec.Command("sh", "-c", cmd).Output()
	return err == nil, string(output)
}

var key_pressed = false

func print_line(text string, keypress chan []byte, echo_state bool) {
	var counter = 0
	var b []byte = make([]byte, 1)
	for _, char := range text {
		print(string(char))
		if echo_state == false {
			counter++
			go func() { os.Stdin.Read(b); keypress <- b }()
			select {
			case key := <-keypress:
				// only skip if enter or space was pressed
				if key[0] == 10 || key[0] == 32 {
					key_pressed = true
					fmt.Println(text[counter:len(text)])
					return
				}
			default:
				fmt.Print("")
			}
		}
		time.Sleep(50 * time.Millisecond)
		if char == '.' || char == '!' || char == '?' {
			time.Sleep(500 * time.Millisecond)
		}
	}
	print("\n")
}

func PrintText(text string, pretty_print bool) {
	var echo_state bool = true
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		if echo_state == false {
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		}
		os.Exit(0)
	}()

	if pretty_print == false {
		// disable input buffering
		err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		if err == nil {
			// do not display entered characters on the screen
			exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
			echo_state = false
			// restore the echoing state when exiting
			defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		}
	}

	keypress := make(chan []byte, 1)
	lines := strings.Split(text, "\n")
	var counter = 0
	for _, line := range lines {
		counter += len(line) + 1
		if counter > len(text) {
			counter = len(text)
		}
		print_line(line, keypress, echo_state)
		if key_pressed {
			fmt.Println(text[counter:len(text)])
			break
		}
	}
}

func MarkdownToTerminal(text string) string {
	bold_regex, _ := regexp.Compile(`\*\*([^\*]+)\*\*`)
	italic_regex, _ := regexp.Compile(`\*([^\*]+)\*`)
	header_regex, _ := regexp.Compile(`^\s*\#+\s*(.+)`)

	bold := "\033[1m"
	inverse := "\033[7m"
	underline_bold := "\033[1;4m"
	reset := "\033[0m"

	text = regexReplaceFunc(bold_regex, text, bold, reset)
	text = regexReplaceFunc(italic_regex, text, inverse, reset)
	text = regexReplaceFunc(header_regex, text, underline_bold, reset)

	return text
}

func regexReplaceFunc(r *regexp.Regexp, text string, start string, end string) string {
	return r.ReplaceAllStringFunc(text, func(s string) string {
		return start + r.FindStringSubmatch(s)[1] + end
	})
}

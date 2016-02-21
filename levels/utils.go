package levels

import (
	"os/exec"
	"regexp"
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

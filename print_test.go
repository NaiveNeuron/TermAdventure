package main

import (
	"./levels"
)

// Test that tries to simulate condition with "last line counter overflow"
// where if user tried to skip text on last line of level description
// counter would exceed length of text and panic with index out of bounds
// would follow.
func ExampleLastLineCounterOverflow() {
	keypress := make(chan []byte, 1)
	markdown_text := levels.MarkdownToTerminal("testing sequence for last" +
		" line counter overflow")
	go func() {
		var space = []byte{10}
		keypress <- space
	}()
	levels.PrettyPrintText(markdown_text, keypress, false)
	// Output:
	// testing sequence for last line counter overflow
}

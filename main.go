package main

import (
	"./levels"
)

func main() {
	challenge := levels.NewChallenge("first")
	challenge.AddLevel(levels.Level{
		TestCmd: `true`,
		Text: `
# First alternative bash excercise

Hi there.

I understand that a command line talking to you might not make sense right now
but bear with me, things will get better.

First of all, you need to learn how to move from one directory to another. This
can be done using the "cd" commad which is a very short abbreviation for
'change directory'.  For instance if I wanted to go to the /usr directory
I would type

	$ cd /usr

Got it?

I am pretty sure you do. Now use this cd comand and go to the /tmp directory
`,
	})

	challenge.AddLevel(levels.Level{
		Name:    `l01`,
		TestCmd: `[[ $(pwd) == "/tmp" ]]`,
		Text: `
I see you made it, awesome !

Now try to get back to your home directory. This is the directory you are in
when you log into your account or open a new shell. It is pretty tough to
remember its particular name so the old UNIX hackers have simplified it for
you: if you execute the "cd" command without any additional parameters, you
will end up in your home directory, wherever that might be.
`,
	})

	challenge.AddLevel(levels.Level{
		Name:    `l02`,
		TestCmd: `[[ $(pwd) == "$HOME" ]]`,
		Text: `
I see you made it again, awesome! That's all for now, so lay back and enjoy your
shell!
`,
	})

	challenge.AddLevel(levels.Level{
		Name:    `l03`,
		TestCmd: `false`,
		Text:    ``,
	})

	challenge.LoadCfg()

	if challenge.CheckCurrentLevel() {
		challenge.PrintCurrentLevel()
		challenge.IncreaseLevel()
	}

	challenge.PrintIdentifier()
}

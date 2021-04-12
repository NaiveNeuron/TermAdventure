name: l00
test: test $(pwd) = "/tmp"
next: [{{ generate_levels "l01-" .names "%s%d" }}]

# First alternative bash excercise

Hi there.

I understand that a command line talking to you might not make sense right now
but bear with me, things *will* get more interesting.

First of all, you need to learn how to move from one directory to another. This
can be done using the "cd" command which is a very short abbreviation for
**'change directory'**.  For instance if I wanted to go to the /usr directory
I would type

	$ cd /usr

Got it?

I am pretty sure you do. Now use this cd command and go to the /tmp directory

--------------------
{{range $index, $name := .names}}
name: l01-{{add $index 1}}
precmd: echo "Hi there!" > /tmp/ta_test
postcmd: rm /tmp/ta_test
test: test $(pwd) = "$HOME"
next: [l02]

I see you made it to {{$name}}, awesome !

Now try to get back to your home directory. This is the directory you are in
when you log into your account or open a new shell. It is pretty tough to
remember its particular name so the old UNIX hackers have simplified it for
you: if you execute the "cd" command without any additional parameters, you
will end up in your home directory, wherever that might be.

(Note: you might want to check /tmp/ta_test)

--------------------
{{end}}
name: l02
test: false

I see you made it again, awesome! That's all for now, so lay back and enjoy your
shell!

--------------------


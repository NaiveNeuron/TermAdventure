# go-term-adventure

A go library for creating good old text adventures in and for the *nix terminals

In order to run the example runner (`sample_challenge.sh`) you need to clone
this repository, install the Go language binary and run

        go build -ldflags "-X main.encryption_key=example_key_1234"

in this repository.

The respective levels (or stages) of your challenges (shell sessions the
runners start) are described in `.gta` files. Levels are separated by Markdown
horizontal lines with at least 10 dashes.

In order for a level to be recognized, it needs to have two parts: `metadata``
and `text`. As expected, the `text` is just Markdown formatted text that gets
printed in terminal and should guide the user. The `metadata` consist of a YAML
document that specifies metadata for a given level. The `metadata` are
separated from `text` by two new lines. An example challenge (.gta) file might
therefore look as follows:

```
test: true

In this level your task is to change your working directory to '/tmp'.

--------------------

name: level1
test: test $(pwd) = "/tmp"

Awesome, you made it to /tmp. Now get back to your home directory.

--------------------

name: level2
test: test $(pwd) = "$HOME"

I see you made it again, awesome! That's all for now, so lay back and enjoy your
shell!

--------------------

name: finish
test: false

--------------------

```

Note that it is not possible to get to the `finish` level, since the `false`
command never produces exit status code `0`, which is necessary for the test to
pass. Also note that the first level uses the test command `true` which always
produces exit status code `0`, and therefore the text of this level will be
printed and the user can proceed to `level1`. In order to solve that level the
`test` command needs to pass (return exit status code `0`). Once it does, the
same procedure takes place.

The generated binary (`go-term-adventure`) also provides an option to check,
whether the `.gta` file is loaded correctly. In order to check file
`sample.gta` you can run

        $ ./go-term-adventure --print sample.gta

Quickstart
==========


Running the demo
----------------

In order to run the example runner (`sample_challenge.sh`) you need to clone
this repository, install the Go language binary and run::

        go build -ldflags "-X main.encryption_key=example_key_1234"


Structure of a GTA level
------------------------

The respective levels (or stages) of your challenges (shell sessions the
runners start) are described in ``.gta`` files. Levels are separated by Markdown
horizontal lines with at least 10 dashes.

In order for a level to be recognized, it needs to have two parts: ``metadata``
and ``text``. As expected, the ``text`` is just Markdown formatted text that gets
printed in terminal and should guide the user. The ``metadata`` consist of a YAML
document that specifies metadata for a given level. The ``metadata`` are
separated from ``text`` by two new lines. An example challenge (.gta) file might
therefore look as follows::

    name: level0
    test: test $(pwd) = "/tmp"

    In this level your task is to change your working directory to '/tmp'.

    --------------------

    name: level1
    test: test $(pwd) = "$HOME"

    Awesome, you made it to /tmp. Now get back to your home directory.

    --------------------

    name: level2
    test: false

    I see you made it again, awesome! That's all for now, so lay back and enjoy your
    shell!


The basic operation can be described follows:

1. User is first 'dropped' into the first level.
2. The level's text is printed.
3. After each command the ``test`` of the current level is executed.
4. If the ``test`` passes (that is the exit status code of the command
   mentioned in the ``test`` field is ``0``), user's level gets changed to
   the next level and the application proceeds with 2.

Note that in our case it is not possible to get out of ``level2``, since
the ``false`` command never produces an exit status code ``0``, which is
necessary for the test to pass.

Validating a given ``.gta`` file
--------------------------------

The generated binary (``go-term-adventure``) also provides an option to check,
whether the ``.gta`` file is loaded correctly. In order to check file
``sample.gta`` you can run::

        $ ./go-term-adventure --print sample.gta


Validating user's progress
--------------------------

The binary uses a hash of many interesting things to keep track of the level
you are currently in. Depending on the name of the challenge and whether you
use the attached ``challenger.sh``, it can be found in the following file::

    $HOME/$CHALLENGE/config.ini


Once you have this hash and the home directory of the user which has managed to
enter the level denoted by this hash, you can reverse-lookup the name of that
level by executing::

    $ ./go-term-adventure --detect-level ./sample.gta $HASH $HOMEDIR


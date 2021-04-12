#!/bin/bash
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TA_BIN=${TA_BIN=:"$CURRENT_DIR/term-adventure"}
if [ ! -e $CHALLENGE_FILE ];
then
    echo "Challenge file $CHALLENGE_FILE does not exist."
    exit 1;
fi
CHALLENGE_NAME=${CHALLENGE_NAME=:"$(basename $CHALLENGE_FILE | sed 's/\.[^\.]*$//')"}

# Prepare global ENV variables for child processes
export TA_BIN
export CHALLENGE_FILE

# Run bash
bash --rcfile $CURRENT_DIR/ta_bashrc

# Clean up used files
rm -rf $HOME/.tahistory
rm -rf $HOME/.config/$CHALLENGE_NAME

# Unset global ENV variables
unset TA_BIN
unset CHALLENGE_FILE

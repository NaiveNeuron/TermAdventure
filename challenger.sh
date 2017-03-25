#!/bin/bash
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GTA_BIN=${GTA_BIN=:"$CURRENT_DIR/go-term-adventure"}
if [ ! -e $CHALLENGE_FILE ];
then
    echo "Challenge file $CHALLENGE_FILE does not exist."
    exit 1;
fi
CHALLENGE_NAME=${CHALLENGE_NAME=:"$(basename $CHALLENGE_FILE | sed 's/\.[^\.]*$//')"}

# Prepare global ENV variables for child processes
export GTA_BIN
export CHALLENGE_FILE

# Run bash
bash --rcfile $CURRENT_DIR/gta_bashrc

# Clean up used files
rm -rf $HOME/.gtahistory
rm -rf $HOME/.config/$CHALLENGE_NAME

# Unset global ENV variables
unset GTA_BIN
unset CHALLENGE_FILE

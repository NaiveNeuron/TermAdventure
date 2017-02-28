#!/bin/bash
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GTA_BIN=${GTA_BIN=:"$CURRENT_DIR/go-term-adventure"}
if [ ! -e $CHALLENGE_FILE ];
then
    echo "Challenge file $CHALLENGE_FILE does not exist."
    exit 1;
fi
CHALLENGE_NAME="$(basename $CHALLENGE_FILE | sed 's/\..*$//')"

PROMPT_COMMAND="history -a; $GTA_BIN $CHALLENGE_FILE"\
        bash --rcfile $CURRENT_DIR/gta_bashrc | tee $HOME/.gtaoutput

rm -rf $HOME/.gtahistory
rm -rf $HOME/.config/$CHALLENGE_NAME

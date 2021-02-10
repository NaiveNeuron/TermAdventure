#!/bin/bash
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
CHALLENGE_FILE="$CURRENT_DIR/sample_challenge.ta.enc"
TA_BIN="$CURRENT_DIR/term-adventure"

source $CURRENT_DIR/challenger.sh

#!/bin/bash

PROMPT_COMMAND="$(pwd)/go-term-adventure $(pwd)/sample_challenge.gta" HISTFILE=$HOME/.gtahistory bash
rm -rf $HOME/.gtadvhistory
rm -rf $HOME/.config/sample_challenge

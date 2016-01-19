#!/bin/bash

PROMPT_COMMAND=$(pwd)/go-term-adventure HISTFILE=$HOME/.advhistory bash
rm -rf $HOME/.advhistory
rm -rf $HOME/.config/first

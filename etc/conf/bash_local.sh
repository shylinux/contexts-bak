#!/bin/bash

echo "" > ~/.hushlogin
export BASH_SILENCE_DEPRECATION_WARNING=1

export LC_ALL=en_US.utf-8
export CTX_ROOT=${CTX_ROOT:=~/contexts}

ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias t "tmux attach"

ish_sys_path_load
ish_sys_path_insert $PWD/usr/publish
ish_sys_path_insert $PWD/bin

ps aux |grep -v grep |grep ice.bin &>/dev/null && return
ps aux |grep -v grep |grep tmux &>/dev/null && return
cd $CTX_ROOT && source etc/miss.sh

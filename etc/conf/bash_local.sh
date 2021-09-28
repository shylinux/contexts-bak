#!/bin/bash

export LC_ALL=en_US.utf-8
export CTX_ROOT=${CTX_ROOT:=~/contexts}

ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias t "tmux attach"
ish_sys_path_insert $CTX_ROOT/usr/publish
ish_sys_path_insert $CTX_ROOT/usr/install/vim81/_install/bin

ps aux |grep -v grep |grep ice.bin &>/dev/null && return
ps aux |grep -v grep |grep tmux &>/dev/null && return
cd $CTX_ROOT && source etc/miss.sh

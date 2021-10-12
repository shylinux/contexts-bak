#!/bin/bash

export LC_ALL=en_US.utf-8
export CTX_ROOT=${CTX_ROOT:=~/contexts}

ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias t "tmux attach"
ish_sys_path_insert $CTX_ROOT/usr/publish

touch $CTX_ROOT/etc/path
for line in `cat $CTX_ROOT/etc/path`; do
    ish_sys_path_insert $line
    ish_log_debug "path" $line
done

ps aux |grep -v grep |grep ice.bin &>/dev/null && return
ps aux |grep -v grep |grep tmux &>/dev/null && return
cd $CTX_ROOT && source etc/miss.sh

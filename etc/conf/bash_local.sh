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
return

ps aux |grep -v grep |grep ice.bin &>/dev/null && return
ps aux |grep -v grep |grep tmux &>/dev/null && return
sleep 1 && cd $CTX_ROOT && source etc/miss.sh
return

# shy
su - shy -c "cd /home/shy/contexts/ && ./bin/ice.bin forever serve &"
cd /home/shy/contexts/usr/local/daemon/10000 && ./sbin/nginx -p $PWD

# dev
su - harveyshao -c "cd /home/harveyshao/contexts/ && ./bin/ice.bin forever serve dev dev &"
cd /home/harveyshao/contexts/usr/local/daemon/10002 && ./sbin/nginx -p $PWD

# mac
cd /Users/harveyshao/contexts/ && ./bin/ice.bin forever serve dev shy &


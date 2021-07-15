#!/bin/bash

export LC_ALL=en_US.utf-8
export CTX_ROOT=${CTX_ROOT:=~/contexts}

export PATH=$CTX_ROOT/usr/install/vim81/_install/bin:$PATH
alias vi=vim

ps aux |grep ice.bin &>/dev/null && return
ps aux |grep tmux &>/dev/null && return
cd $CTX_ROOT && source etc/miss.sh

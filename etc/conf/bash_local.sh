#!/bin/bash

# export LC_ALL=zh_CN.utf-8
export LC_ALL=en_US.utf-8

export CTX_ROOT=${CTX_ROOT:=~/contexts}
export GOROOT=$CTX_ROOT/usr/local/go

export PATH=$CTX_ROOT/bin:$CTX_ROOT/usr/publish:$CTX_ROOT/usr/local/bin:$GOROOT/bin:$GOBIN:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin

if [ -f ~/.ish/plug.sh ] && source ~/.ish/plug.sh; then
    require conf.sh

    ish_ctx_cli_prompt
fi

ps aux |grep tmux &>/dev/null && return
ps aux |grep ice.bin &>/dev/null && return
cd $CTX_ROOT && source etc/miss.sh
[ -f ~/.bash_local ] && source ~/.bash_local

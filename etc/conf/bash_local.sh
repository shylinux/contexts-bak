#!/bin/bash

export CTX_ROOT=${CTX_ROOT:=~/contexts}
export GOROOT=$CTX_ROOT/usr/local/go

export PATH=$CTX_ROOT/bin:$CTX_ROOT/usr/local/bin:$CTX_ROOT/usr/local/go/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin

if [ -f ~/.ish/plug.sh ] && source ~/.ish/plug.sh; then
    require conf.sh

    ish_ctx_cli_prompt
fi


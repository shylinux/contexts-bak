#!/bin/bash

[ -f $PWD/.ish/plug.sh ] || [ -f $HOME/.ish/plug.sh ] || git clone ${ISH_CONF_HUB_PROXY:="https://"}shylinux.com/x/intshell $PWD/.ish
if [ "$ISH_CONF_PRE" = "" ]; then
    source $PWD/.ish/plug.sh || source $HOME/.ish/plug.sh
fi

require sys/cli/file.sh
ish_sys_path_load

[ -f ~/.bash_profile ] || echo "source ~/.bashrc" >> ~/.bash_profile
ish_sys_link_create ~/.bash_local $PWD/etc/conf/bash_local.sh

require miss.sh
ish_miss_prepare_compile
ish_miss_prepare_develop
ish_miss_prepare_install

# ish_miss_prepare word-dict
# ish_miss_prepare wubi-dict

ish_miss_prepare linux-story
ish_miss_prepare nginx-story
ish_miss_prepare golang-story
ish_miss_prepare redis-story
ish_miss_prepare mysql-story
ish_miss_prepare release

ish_miss_prepare_contexts
ish_miss_prepare_intshell
ish_miss_prepare_icebergs
ish_miss_prepare_toolkits
ish_miss_prepare_volcanos
ish_miss_prepare_learning

ish_miss_make

ish_sys_link_create ~/.vim_local.vim $PWD/etc/conf/vim_local.vim
require dev/vim/vim.sh
ish_dev_vim_prepare

require dev/tmux/tmux.sh
ish_dev_tmux_prepare

ish_miss_prepare_session miss miss


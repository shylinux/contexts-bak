#!/bin/bash
if [ "$ISH_CONF_PRE" = "" ]; then
    [ -f $PWD/.ish/plug.sh ] || [ -f $HOME/.ish/plug.sh ] || git clone ${ISH_CONF_HUB_PROXY:="https://"}github.com/shylinux/intshell $PWD/.ish
    source $PWD/.ish/plug.sh || source $HOME/.ish/plug.sh
fi; [ "$PWD" = "$HOME" ] || ish_miss_create_link $HOME/.ish $PWD/.ish

[ -f ~/.bash_profile ] || echo "source ~/.bashrc" >> ~/.bash_profile
ish_miss_create_link ~/.bash_local $PWD/etc/conf/bash_local.sh
require dev/git/git.sh
ish_dev_git_prepare

require miss.sh
ish_miss_prepare_compile
ish_miss_prepare_develop
ish_miss_prepare_install

ish_miss_prepare_contexts
ish_miss_prepare_intshell
ish_miss_prepare_icebergs
ish_miss_prepare_toolkits
ish_miss_prepare_volcanos
ish_miss_prepare_learning

# ish_miss_prepare wubi-dict
# ish_miss_prepare word-dict

ish_miss_prepare linux-story
ish_miss_prepare nginx-story
ish_miss_prepare golang-story
ish_miss_prepare redis-story
ish_miss_prepare mysql-story

make

ish_miss_create_link ~/.vim_local.vim $PWD/etc/conf/vim_local.vim
require dev/vim/vim.sh
ish_dev_vim_prepare

require dev/tmux/tmux.sh
ish_dev_tmux_prepare

ish_miss_prepare_session miss miss


#!/bin/bash

[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone ${ISH_CONF_HUB_PROXY:="https://"}github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh
require miss.sh

ish_miss_prepare_compile
ish_miss_prepare_install

ish_miss_prepare_volcanos
ish_miss_prepare_learning
ish_miss_prepare_icebergs
ish_miss_prepare_toolkits
ish_miss_prepare_intshell
ish_miss_prepare_contexts

# ish_miss_prepare wubi-dict
# ish_miss_prepare word-dict

ish_miss_prepare linux-story
ish_miss_prepare nginx-story
ish_miss_prepare golang-story
ish_miss_prepare redis-story
ish_miss_prepare mysql-story

ish_miss_prepare_bash() {
    ish_miss_create_link ~/.bashrc $PWD/etc/conf/bashrc
    ish_miss_create_link ~/.zshrc $PWD/etc/conf/zshrc
    ish_miss_create_link ~/.ish $PWD/.ish

    if [ -z `cat ~/.bashrc |grep "source ~/.ish/plug.sh"` ]; then
        if [ -z `cat ~/.bashrc |grep "source $PWD/etc/conf/bashrc"` ]; then
            echo -e "\nsource $PWD/etc/conf/bashrc\n" >> ~/.bashrc
        fi
    fi

    if [ -z `cat ~/.zshrc |grep "source ~/.ish/plug.sh"` ]; then
        if [ -z `cat ~/.zshrc |grep "source $PWD/etc/conf/zshrc"` ]; then
            echo -e "\nsource $PWD/etc/conf/zshrc\n" >> ~/.zshrc
        fi
    fi
}
# ish_miss_prepare_bash

require misc/tmux/tmux.sh
require misc/git/git.sh
require misc/vim/vim.sh

ish_ctx_dev_tmux_prepare
ish_ctx_dev_git_prepare
ish_ctx_dev_vim_prepare

# ish_miss_prepare_develop
make
ish_miss_prepare_session miss

#!/bin/bash
[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone https://github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh
require miss.sh

ish_miss_prepare_compile
ish_miss_prepare_install

ish_miss_prepare_volcanos
ish_miss_prepare learning
ish_miss_prepare_icebergs
ish_miss_prepare toolkits
ish_miss_prepare_intshell
ish_miss_prepare_contexts

# ish_miss_prepare wubi-dict
# ish_miss_prepare word-dict

ish_miss_prepare linux-story
ish_miss_prepare nginx-story
ish_miss_prepare golang-story
ish_miss_prepare redis-story
ish_miss_prepare mysql-story

ish_miss_prepare_vim() {
    ish_miss_create_link ~/.vimrc $PWD/etc/conf/vimrc
    ish_miss_create_link ~/.auto.vim $PWD/usr/publish/auto.vim

    mkdir -p ~/.vim/autoload
    ish_miss_create_link ~/.vim/autoload/plug.vim $PWD/etc/conf/plug.vim

    mkdir -p ~/.vim/syntax
    ish_miss_create_link ~/.vim/syntax/sh.vim $PWD/etc/conf/sh.vim
    ish_miss_create_link ~/.vim/syntax/shy.vim $PWD/etc/conf/shy.vim
    ish_miss_create_link ~/.vim/syntax/go.vim $PWD/etc/conf/go.vim
    ish_miss_create_link ~/.vim/syntax/javascript.vim $PWD/etc/conf/javascript.vim
}
ish_miss_prepare_git() {
    git config --global alias.s status
    git config --global alias.b branch
    git config --global credential.helper store
}
ish_miss_prepare_tmux() {
    ish_miss_create_link ~/.tmux.conf $PWD/etc/conf/tmux.conf
}
ish_miss_prepare_bash() {
    ish_miss_create_link ~/.bashrc $PWD/etc/conf/bashrc
    ish_miss_create_link ~/.ish $PWD/.ish
}

ish_miss_prepare_bash
ish_miss_prepare_tmux
ish_miss_prepare_git
ish_miss_prepare_vim

# ish_miss_prepare_develop
make
ish_miss_prepare_session miss
export PATH=$PWD/bin:$PWD/usr/local/bin:$PWD/usr/local/go/bin:$PATH


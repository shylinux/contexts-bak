#!/bin/sh

# 加载框架
[ -d ~/.ish ] || [ -d usr/intshell ] || git clone https://github.com/shylinux/intshell usr/intshell
[ "$ISH_CONF_PRE" != "" ] || source usr/intshell/plug.sh || source ~/.ish/plug.sh
declare -f ish_help_repos || require conf.sh

ish_miss_pwd=${PWD}
ish_miss_pwd() { echo $ish_miss_pwd; }
ish_miss_reload() { require ${ish_miss_pwd}/$BASH_SOURCE; }

ish_miss_prepare_syntax_source=${PWD}/conf
ish_miss_prepare_syntax_target=~/.vim/syntax
ish_miss_prepare_syntax_language="sh go shy javascript"
ish_miss_prepare_syntax() { local prefix=ish_miss_prepare_syntax
    local source=$(ish_get $prefix source)
    local target=$(ish_get $prefix target) && [ -d $target ] || mkdir $target
    for _l in $(ish_get $prefix language); do local file=$target/$_l.vim
        [ -f $file ] || ln $source/$_l.vim $file
    done
    ish_show -green "vim syntax script:" && ls -lht $target
}
ish_miss_prepare() {
    ish_miss_prepare_syntax
}
return

create() {
    tmux new-session -d -s miss -n shy
    tmux split-window -p 30 -t miss:shy.1
    tmux split-window -h -t miss:shy.2
    tmux send-keys -t miss:shy.3 "tail -f bin/boot.log" Enter
    tmux send-keys -t miss:shy.2 "bin/ice.sh start serve shy" Enter
    tmux send-keys -t miss:shy.1 "vim" Enter
}

tmux has-session -t miss &>/dev/null || create
tmux attach-session -t miss


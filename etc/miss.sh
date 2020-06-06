[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone https://github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh
# declare -f ish_help_repos &>/dev/null || require conf.sh

require show.sh
require help.sh
require miss.sh
# cd $ISH_CONF_TASK

ish_miss_prepare_compile
ish_miss_prepare_install
ish_miss_prepare_session miss

ish_miss_prepare_volcanos
ish_miss_prepare_icebergs
ish_miss_prepare toolkits
# ish_miss_prepare_intshell
# ish_miss_prepare learning

# ish_miss_prepare wubi-dict
# ish_miss_prepare word-dict

repos=(volcanos icebergs intshell contexts toolkits learning)
ish_miss_pull() {
    for p in $repos; do
        cd usr/$p && echo && ish_show -g $PWD
        git pull
        cd -
    done
}
ish_miss_status() {
    for p in $repos; do
        cd usr/$p && echo && ish_show -g $PWD
        git status
        cd -
    done
}
ish_miss_build() {
    local target=bin/ice.bin source=src/main.go
    ish_log_debug build $source
	go build -v -o $target $source && chmod u+x $target && ./bin/ice.sh restart
    ish_log_debug build $target
}
ish_miss_build_windows() {
    local target=usr/publish/ice.windows.amd64 source=src/main.go
    ish_log_debug build $source
    GOOS=windows go build -v -o $target $source && chmod u+x $target
    ish_log_debug build $target
}

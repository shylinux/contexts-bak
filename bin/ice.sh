#! /bin/sh

export PATH=${PWD}/bin:${PWD}:$PATH
export ctx_dev=${ctx_dev:="https://shylinux.com:443"}
export ctx_mod=${ctx_mod:="gdb,log,ssh,ctx"}
export ctx_pid=${ctx_pid:=var/run/ice.pid}
export ctx_log=${ctx_log:=bin/boot.log}

prepare() {
    [ -d bin ] || mkdir bin
    [ -e bin/ice.sh ] || curl -sq $ctx_dev/publish/ice.sh -o bin/ice.sh && chmod u+x bin/ice.sh
    [ -e bin/ice.bin ] && chmod u+x bin/ice.bin && return

    bin="ice"
    case `uname -s` in
        Darwin) bin=${bin}.darwin ;;
        Linux) bin=${bin}.linux ;;
        *) bin=${bin}.windows ;;
    esac
    case `uname -m` in
        x86_64) bin=${bin}.amd64 ;;
        arm*) bin=${bin}.arm ;;
        *) bin=${bin}.386 ;;
    esac
    curl -sq $ctx_dev/publish/${bin} -o bin/ice.bin && chmod u+x bin/ice.bin
 }
start() {
    trap HUP hup && while true; do
        date && echo -e "\n\nrestarting..."
        echo -e "ctx_dev: $ctx_dev ctx_mod: $ctx_mod ctx_cmd: $ctx_cmd"
        ice.bin $@ 2>$ctx_log && break
    done
}
restart() {
    [ -e $ctx_pid ] && kill -2 `cat $ctx_pid` || echo
}
shutdown() {
    [ -e $ctx_pid ] && kill -3 `cat $ctx_pid` || echo
}
serve() {
    prepare && shutdown && start $@
}

cmd=$1 && [ -n "$cmd" ] && shift || cmd="serve serve dev"; $cmd $*

#! /bin/sh

export PATH=${PWD}:${PWD}/bin:$PATH
export ctx_pid=${ctx_pid:=var/run/ice.pid}
export ctx_log=${ctx_log:=bin/boot.log}

prepare() {
    [ -d bin ] || mkdir bin
    [ -e bin/ice.sh ] || curl $ctx_dev/publish/ice.sh -o bin/ice.sh && chmod u+x bin/ice.sh
    [ -e bin/ice.bin ] && chmod u+x bin/ice.bin && return

    bin="ice"
    case `uname -s` in
        Darwin) bin=${bin}.darwin ;;
        Linux) bin=${bin}.linux ;;
        *) bin=${bin}.windows ;;
    esac
    case `uname -m` in
        x86_64) bin=${bin}.amd64 ;;
        i686) bin=${bin}.386 ;;
        arm*) bin=${bin}.arm ;;
    esac
    curl $ctx_dev/publish/${bin} -o bin/ice.bin && chmod u+x bin/ice.bin
 }
start() {
    trap HUP hup && while true; do
        date && ice.bin $@ 2>$ctx_log && echo -e "\n\nrestarting..." || break
    done
}
serve() {
    prepare && shutdown && start $@
}
restart() {
    [ -e $ctx_pid ] && kill -2 `cat $ctx_pid` || echo
}
shutdown() {
    [ -e $ctx_pid ] && kill -3 `cat $ctx_pid` || echo
}

cmd=$1 && [ -n "$cmd" ] && shift || cmd=serve
$cmd $*

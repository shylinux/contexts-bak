#! /bin/sh

export ctx_log=${ctx_log:=bin/boot.log}
export ctx_pid=${ctx_pid:=var/run/ice.pid}
export ctx_mod=${ctx_mod:=gdb,log,ssh,ctx}

prepare() {
    [ -d bin ] || mkdir bin
    [ -e bin/ice.sh ] || (curl -o bin/ice.sh -fsSL $ctx_dev/publish/ice.sh || wget -O bin/ice.sh $ctx_dev/publish/ice.sh) && chmod u+x bin/ice.sh
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
    (curl -o bin/ice.bin -fsSL $ctx_dev/publish/${bin} || wget -O bin/ice.bin $ctx_dev/publish/${bin}) && chmod u+x bin/ice.bin
}

restart() {
    [ -e $ctx_pid ] && kill -2 `cat $ctx_pid` &>/dev/null || echo
}
start() {
    prepare
    trap HUP hup && while true; do
        date && bin/ice.bin $@ 2>$ctx_log && echo -e \"\n\nrestarting...\" && break
    done
}
stop() {
    [ -e $ctx_pid ] && kill -3 `cat $ctx_pid` &>/dev/null || echo
}
serve() {
    stop && start $@
}

cmd=$1 && [ -n \"$cmd\" ] && shift || cmd=serve
$cmd $*

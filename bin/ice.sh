#! /bin/sh

export ctx_log=${ctx_log:=bin/boot.log}
export ctx_pid=${ctx_pid:=var/run/ice.pid}
export ctx_mod=${ctx_mod:=gdb,log,ssh,ctx}

restart() {
    [ -e $ctx_pid ] && kill -2 `cat $ctx_pid` &>/dev/null || echo
}
start() {
    trap HUP hup && while true; do
        date && ice.bin $@ 2>$ctx_log && echo -e \"\n\nrestarting...\" && break
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

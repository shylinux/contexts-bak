#! /bin/sh

export ctx_log=${ctx_log:=bin/boot.log}
export ctx_pid=${ctx_pid:=var/run/ice.pid}


start() {
    trap HUP hup && while true; do
        date && bin/ice.bin $@ 2>$ctx_log && break || echo -e \"\n\nrestarting...\"
    done
}
restart() {
    [ -e $ctx_pid ] && kill -2 `cat $ctx_pid` &>/dev/null || echo
}
stop() {
    [ -e $ctx_pid ] && kill -3 `cat $ctx_pid` &>/dev/null || echo
}
serve() {
    stop && start $@
}

cmd=$1 && [ -n \"$cmd\" ] && shift || cmd=serve
$cmd $*

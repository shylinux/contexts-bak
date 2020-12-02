package main

import (
	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"
	_ "github.com/shylinux/icebergs/misc/alpha"
	_ "github.com/shylinux/icebergs/misc/chrome"
	_ "github.com/shylinux/icebergs/misc/input"

	_ "github.com/shylinux/icebergs/misc/lark"
	_ "github.com/shylinux/icebergs/misc/mp"
	_ "github.com/shylinux/icebergs/misc/wx"

	_ "github.com/shylinux/golang-story/src/compile"
	_ "github.com/shylinux/golang-story/src/project"
	_ "github.com/shylinux/golang-story/src/runtime"
	_ "github.com/shylinux/nginx-story/src/server"
	_ "github.com/shylinux/redis-story/src/client"
	_ "github.com/shylinux/redis-story/src/server"

	_ "github.com/shylinux/linux-story/iso/centos"
	_ "github.com/shylinux/linux-story/src/busybox"
	_ "github.com/shylinux/linux-story/src/gcc"
	_ "github.com/shylinux/linux-story/src/gdb"
	_ "github.com/shylinux/linux-story/src/glibc"
	_ "github.com/shylinux/linux-story/src/kernel"
	_ "github.com/shylinux/linux-story/src/qemu"
	_ "github.com/shylinux/mysql-story/src/client"
	_ "github.com/shylinux/mysql-story/src/search"
	_ "github.com/shylinux/mysql-story/src/server"
)

func main() { print(ice.Run()) }

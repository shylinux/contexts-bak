package main

import (
	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/icebergs/misc/alpha"
	_ "github.com/shylinux/icebergs/misc/chrome"
	_ "github.com/shylinux/icebergs/misc/input"
	_ "github.com/shylinux/icebergs/misc/totp"

	_ "github.com/shylinux/icebergs/misc/es"
	_ "github.com/shylinux/icebergs/misc/lark"
	_ "github.com/shylinux/icebergs/misc/mp"
	_ "github.com/shylinux/icebergs/misc/wx"

	_ "github.com/shylinux/golang-story/src/compile"
	_ "github.com/shylinux/golang-story/src/project"
	_ "github.com/shylinux/golang-story/src/runtime"
	_ "github.com/shylinux/linux-story/src/gcc"
	_ "github.com/shylinux/linux-story/src/glibc"
	_ "github.com/shylinux/linux-story/src/kernel"
	_ "github.com/shylinux/mysql-story/src/server"
	_ "github.com/shylinux/nginx-story/src/server"
	_ "github.com/shylinux/redis-story/src/client"
	_ "github.com/shylinux/redis-story/src/server"
)

func main() { ice.Run() }

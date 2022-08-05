package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/alpha"
	_ "shylinux.com/x/icebergs/misc/chrome"
	_ "shylinux.com/x/icebergs/misc/coder"
	_ "shylinux.com/x/icebergs/misc/input"

	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/lark"
	_ "shylinux.com/x/icebergs/misc/mp"
	_ "shylinux.com/x/icebergs/misc/node"
	_ "shylinux.com/x/icebergs/misc/wework"
	_ "shylinux.com/x/icebergs/misc/wx"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/grafana"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/prometheus"
	_ "shylinux.com/x/golang-story/src/runtime"

	_ "shylinux.com/x/mysql-story/src/client"
	_ "shylinux.com/x/mysql-story/src/search"
	_ "shylinux.com/x/mysql-story/src/server"
	_ "shylinux.com/x/nginx-story/src/server"
	_ "shylinux.com/x/redis-story/src/client"
	_ "shylinux.com/x/redis-story/src/server"
)

func main() { print(ice.Run()) }

package main

import (
	"github.com/shylinux/icebergs"

	_ "github.com/shylinux/icebergs/base/aaa"
	_ "github.com/shylinux/icebergs/base/cli"
	_ "github.com/shylinux/icebergs/base/ctx"
	_ "github.com/shylinux/icebergs/base/web"

	_ "github.com/shylinux/icebergs/base/gdb"
	_ "github.com/shylinux/icebergs/base/lex"
	_ "github.com/shylinux/icebergs/base/log"
	_ "github.com/shylinux/icebergs/base/yac"

	_ "github.com/shylinux/icebergs/base/mdb"
	_ "github.com/shylinux/icebergs/base/nfs"
	_ "github.com/shylinux/icebergs/base/ssh"
	_ "github.com/shylinux/icebergs/base/tcp"

	_ "github.com/shylinux/icebergs/core/chat"
	_ "github.com/shylinux/icebergs/core/code"
	_ "github.com/shylinux/icebergs/core/mall"
	_ "github.com/shylinux/icebergs/core/team"
	_ "github.com/shylinux/icebergs/core/wiki"

	_ "github.com/shylinux/icebergs/misc/docker"
	_ "github.com/shylinux/icebergs/misc/git"
	_ "github.com/shylinux/icebergs/misc/tmux"
	_ "github.com/shylinux/icebergs/misc/vim"
	_ "github.com/shylinux/icebergs/misc/zsh"

	_ "github.com/shylinux/icebergs/misc/alpha"
	_ "github.com/shylinux/icebergs/misc/chrome"
	_ "github.com/shylinux/icebergs/misc/input"
	_ "github.com/shylinux/icebergs/misc/lark"
	// _ "github.com/shylinux/icebergs/misc/md"
	// _ "github.com/shylinux/icebergs/misc/mp"
	// _ "github.com/shylinux/icebergs/misc/pi"
	// _ "github.com/shylinux/icebergs/misc/wx"

	_ "github.com/shylinux/icebergs/misc/totp"
)

func main() { println(ice.Run()) }

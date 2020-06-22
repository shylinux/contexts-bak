package main

import (
	ice "github.com/shylinux/icebergs"

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

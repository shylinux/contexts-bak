package main

import (
	"github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/icebergs/misc/alpha"
	_ "github.com/shylinux/icebergs/misc/chrome"
	_ "github.com/shylinux/icebergs/misc/lark"
	_ "github.com/shylinux/icebergs/misc/mp"
	_ "github.com/shylinux/icebergs/misc/pi"
	_ "github.com/shylinux/icebergs/misc/wx"

	_ "github.com/shylinux/icebergs/misc/auth"
)

func main() {
	println(ice.Run())
}

package main

import (
	"github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/icebergs/misc/alpha"
	_ "github.com/shylinux/icebergs/misc/chrome"
	_ "github.com/shylinux/icebergs/misc/input"
	_ "github.com/shylinux/icebergs/misc/lark"

	// _ "github.com/shylinux/icebergs/misc/md"
	// _ "github.com/shylinux/icebergs/misc/mp"
	// _ "github.com/shylinux/icebergs/misc/pi"
	// _ "github.com/shylinux/icebergs/misc/wx"

	// _ "github.com/shylinux/icebergs/misc/fyne"
	_ "github.com/shylinux/icebergs/misc/totp"
)

func main() { ice.Run() }

package main

import (
	"shylinux.com/x/ice"
)

func main() {
	ice.App("/admin", "管理", `
libra
	dev 开发
		cli.qrcode
		cli.runtime
`)
	ice.App("/vip", "会员", `
libra
	dev 开发
		cli.runtime
		cli.qrcode
`)
	ice.RunServe("port", "9090")
}

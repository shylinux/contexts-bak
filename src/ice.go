package main

import (
	"github.com/shylinux/ice"
)

func main() {
	ice.App("web.demo", "/tool/", func(p *ice.Page) {
		p.Cmd("nav", &ice.Nav{Home: "./", Prefix: "/tool"})
		p.Cmd("cli.system", ice.Arg("pwd"))
		p.Cmd("hash", &ice.Hash{})
		p.Home = "./"
	})
	ice.RunServe("9090")
}

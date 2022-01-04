package main

import (
	"os"
	"path"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/tcp"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/misc/app"
)

func main() {
	os.Chdir(path.Join(os.Getenv(cli.HOME), "contexts"))
	go ice.Run(web.SERVE, tcp.START)
	time.Sleep(time.Second)
	app.Run()
}

package main

import (
	"os"
	"path"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/misc/app"
)

func main() {
	os.Chdir(path.Join("/Users", os.Getenv("USER"), "contexts"))
	go ice.Run("serve", "start")
	time.Sleep(time.Second)
	app.Run()
}

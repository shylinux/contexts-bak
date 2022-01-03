package main

import (
	"os"
	"path"
	"time"

	"github.com/webview/webview"
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/app"
)

func main() {
	os.Chdir(path.Join("/Users", os.Getenv("USER"), "contexts"))
	go ice.Run("serve", "start", "dev", "shy")
	time.Sleep(time.Second)

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("contexts")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:9020")
	w.Run()
}

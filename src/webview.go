package main

import (
	"os"
	"time"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
	"shylinux.com/x/webview"
)

func main() {
	os.Chdir(kit.HomePath("contexts"))
	go ice.Run("serve", "start")
	time.Sleep(time.Second)

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("contexts")
	w.SetSize(1200, 800, webview.HintNone)
	w.Navigate("http://localhost:9020")
	w.Run()
}

package main

import (
	"os"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/misc/webview"
	kit "shylinux.com/x/toolkits"
)

type view struct{ *webview.WebView }

func main() {
	os.Chdir(kit.HomePath("contexts"))
	go ice.Run("serve", "start")
	time.Sleep(time.Second)

	webview.Run(func(w *webview.WebView) interface{} { return view{w} })
}

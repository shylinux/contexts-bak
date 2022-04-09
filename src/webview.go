package main

import (
	"os"
	"strings"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
	"shylinux.com/x/webview"
)

type view struct{ w webview.WebView }

func (v view) Title(text string)  { v.w.SetTitle(text) }
func (v view) Webview(url string) { v.w.Navigate(url) }
func (v view) Open(url string)    { v.w.Navigate(url) }
func (v view) Terminate()         { v.w.Terminate() }
func (v view) Close() {
	if !menu(v.w) {
		v.Terminate()
	}
}

func menu(w webview.WebView) bool {
	kit.Reflect(view{w: w}, func(name string, value interface{}) { w.Bind(name, value) })
	list := []string{}
	ice.Pulse.Cmd(nfs.CAT, "src/webview.txt", func(ls []string, line string) {
		if len(ls) > 1 {
			list = append(list, kit.Format(`<button onclick=%s()>%s</button>`, ls[0], ls[0]))
			w.Bind(ls[0], func() {
				w.SetSize(1200, 800, webview.HintNone)
				w.Navigate(ls[1])
			})
		}
	})

	if len(list) == 0 {
		return false
	}

	w.SetTitle("contexts")
	w.SetSize(200, 60*len(list), webview.HintNone)
	w.Navigate(kit.Format(`data:text/html,
    <!doctype html>
    <html>
	<head><style>button { font-size:24px; margin:10px; width:-webkit-fill-available; display:block; clear:both; }</style></head>
	<body>%s</body>
    </html>`, strings.Join(list, "")))
	return true
}
func main() {
	os.Chdir(kit.HomePath("contexts"))
	go ice.Run("serve", "start")
	time.Sleep(time.Second)

	w := webview.New(true)
	defer w.Destroy()
	defer w.Run()
	if !menu(w) {
		w.SetSize(1200, 800, webview.HintNone)
		w.Navigate("http://localhost:9020")
	}
}

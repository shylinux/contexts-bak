package main

import (
	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base/ssh"
	log "github.com/shylinux/toolkits/logs"

	"os"
	"path"
)

func main() {
	log.LogDisable = true
	args := []string{"ssh.connect", "open", "authfile", path.Join(os.Getenv("HOME"), ".ssh/", path.Base(os.Args[0])+".json")}
	args = append(args, os.Args[1:]...)
	print(ice.Run(args...))
}

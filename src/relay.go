package main

import (
	"os"
	"path"

	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base/ssh"
	log "github.com/shylinux/toolkits/logs"
)

func main() {
	log.LogDisable = true
	defer func() { recover() }()
	args := []string{"ssh.connect", "open", "authfile", path.Join(os.Getenv("HOME"), ".ssh/", path.Base(os.Args[0])+".json")}
	print(ice.Run(append(args, os.Args[1:]...)...))
}

package main

import (
	"os"
	"path"

	ice "shylinux.com/x/icebergs"
	_ "shylinux.com/x/icebergs/misc/ssh"
	log "shylinux.com/x/toolkits/logs"
)

func main() {
	log.LogDisable = true
	defer func() { recover() }()
	args := []string{"ssh.connect", "open", "authfile", path.Join(os.Getenv("HOME"), ".ssh/", path.Base(os.Args[0])+".json")}
	print(ice.Run(append(args, os.Args[1:]...)...))
}

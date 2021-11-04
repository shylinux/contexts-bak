package main

import (
	"os"
	"path"

	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/ssh"
)

func main() {
	defer func() { recover() }()
	args := []string{"ssh.connect", "open", "authfile", path.Join(os.Getenv("HOME"), ".ssh/", path.Base(os.Args[0])+".json")}
	print(ice.Run(append(args, os.Args[1:]...)...))
}

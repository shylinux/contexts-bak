package main

import (
	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base/ssh"

	"os"
	"path"
)

func main() {
	args := []string{"ssh.connect", "open", "authfile", path.Join(os.Getenv("HOME"), ".ssh/"+path.Base(os.Args[0])+".json")}
	args = append(args, os.Args[1:]...)
	print(ice.Run(args...))
}

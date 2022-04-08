module shylinux.com/x/contexts

go 1.11

require shylinux.com/x/ice v0.8.9

require (
	github.com/kr/pty v1.1.8 // indirect
	github.com/sqweek/dialog v0.0.0-20220227145630-7a1c9e333fcf
	github.com/webview/webview v0.0.0-20220407190137-85a0d5c36d07 // indirect
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	shylinux.com/x/golang-story v0.3.8
	shylinux.com/x/icebergs v1.1.6
	shylinux.com/x/linux-story v0.3.6
	shylinux.com/x/mysql-story v0.3.9
	shylinux.com/x/nginx-story v0.3.9
	shylinux.com/x/redis-story v0.3.9
	shylinux.com/x/webview v0.0.1 // indirect
)

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

replace (
	shylinux.com/x/golang-story => ./usr/golang-story
	shylinux.com/x/linux-story => ./usr/linux-story
	shylinux.com/x/mysql-story => ./usr/mysql-story
	shylinux.com/x/nginx-story => ./usr/nginx-story
	shylinux.com/x/redis-story => ./usr/redis-story
)

module shylinux.com/x/contexts

go 1.11

require (
	shylinux.com/x/golang-story v0.4.3
	shylinux.com/x/linux-story v0.4.1
	shylinux.com/x/mysql-story v0.4.4
	shylinux.com/x/nginx-story v0.4.4
	shylinux.com/x/redis-story v0.4.4
)

require (
	github.com/kr/pty v1.1.8 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	shylinux.com/x/ice v0.9.8
	shylinux.com/x/icebergs v1.2.5
	shylinux.com/x/toolkits v0.6.0
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

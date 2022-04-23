module shylinux.com/x/contexts

go 1.11

require shylinux.com/x/ice v0.9.2

require (
	shylinux.com/x/golang-story v0.4.0
	shylinux.com/x/icebergs v1.1.8
	shylinux.com/x/linux-story v0.3.8
	shylinux.com/x/mysql-story v0.4.1
	shylinux.com/x/nginx-story v0.4.1
	shylinux.com/x/redis-story v0.4.1
	shylinux.com/x/toolkits v0.5.8
	shylinux.com/x/webview v0.0.1
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

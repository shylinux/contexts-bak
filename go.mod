module github.com/shylinux/contexts

go 1.11

replace (
	github.com/shylinux/golang-story => ./usr/golang-story
	github.com/shylinux/icebergs => ./usr/icebergs
	github.com/shylinux/linux-story => ./usr/linux-story
	github.com/shylinux/mysql-story => ./usr/mysql-story
	github.com/shylinux/nginx-story => ./usr/nginx-story
	github.com/shylinux/redis-story => ./usr/redis-story
	github.com/shylinux/toolkits => ./usr/toolkits
)

require (
	github.com/shylinux/golang-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/icebergs v0.3.8
	github.com/shylinux/linux-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/mysql-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/nginx-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/redis-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/toolkits v0.2.6
)

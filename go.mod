module shylinux.com/x/contexts

go 1.11

replace (
	shylinux.com/x/golang-story => ./usr/golang-story
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/linux-story => ./usr/linux-story
	shylinux.com/x/mysql-story => ./usr/mysql-story
	shylinux.com/x/nginx-story => ./usr/nginx-story
	shylinux.com/x/redis-story => ./usr/redis-story
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/golang-story v0.0.0-00010101000000-000000000000
	shylinux.com/x/icebergs v0.4.1
	shylinux.com/x/linux-story v0.0.0-00010101000000-000000000000
	shylinux.com/x/mysql-story v0.0.0-00010101000000-000000000000
	shylinux.com/x/nginx-story v0.0.0-00010101000000-000000000000
	shylinux.com/x/redis-story v0.0.0-00010101000000-000000000000
	shylinux.com/x/toolkits v0.3.0
)

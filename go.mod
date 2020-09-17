module github.com/shylinux/contexts

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1
	github.com/shylinux/golang-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/icebergs v0.2.7
	github.com/shylinux/linux-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/mysql-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/nginx-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/redis-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/toolkits v0.1.8
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
)

replace github.com/shylinux/icebergs => ./usr/icebergs

replace github.com/shylinux/toolkits => ./usr/toolkits

replace github.com/shylinux/linux-story => ./usr/linux-story

replace github.com/shylinux/nginx-story => ./usr/nginx-story

replace github.com/shylinux/golang-story => ./usr/golang-story

replace github.com/shylinux/redis-story => ./usr/redis-story

replace github.com/shylinux/mysql-story => ./usr/mysql-story

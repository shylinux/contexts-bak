module contexts

go 1.13

require (
	fyne.io/fyne v1.3.2 // indirect
	github.com/gomarkdown/markdown v0.0.0-20200609195525-3f9352745725 // indirect
	github.com/shylinux/golang-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/icebergs v0.1.23
	github.com/shylinux/nginx-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/redis-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/toolkits v0.1.7
)

replace github.com/shylinux/icebergs => ./usr/icebergs

replace github.com/shylinux/toolkits => ./usr/toolkits

replace github.com/shylinux/nginx-story => ./usr/nginx-story

replace github.com/shylinux/golang-story => ./usr/golang-story

replace github.com/shylinux/redis-story => ./usr/redis-story

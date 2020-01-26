module miss

go 1.13

require (
	github.com/shylinux/contexts v0.0.0-00010101000000-000000000000 // indirect
	github.com/shylinux/icebergs v0.1.9
	github.com/shylinux/toolkits v0.1.0
)

replace (
	github.com/shylinux/contexts => ./
	github.com/shylinux/icebergs => ../icebergs
	github.com/shylinux/toolkits => ../toolkits
)

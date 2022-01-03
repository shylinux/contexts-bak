export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE=shylinux.com,github.com
# export CGO_ENABLED=0

all:
	@echo && date
	[ -f src/version.go ] || echo "package main" > src/version.go
	go build -v -o bin/ice.bin src/main.go src/version.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart

%: src/%.go
	@echo && date
	go build -v -o usr/publish/$@ src/$@.go && chmod u+x usr/publish/$@

ice:
	cat src/binpack.go|sed 's/package main/package ice/g' > usr/release/binpack.go

macapp:
	 mkdir -p usr/publish/contexts.app/Contents/MacOS
	 go build -o usr/publish/contexts.app/Contents/MacOS/contexts src/app.go src/binpack.go
	 open usr/publish/contexts.app

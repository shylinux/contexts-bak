export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE=shylinux.com,github.com
export CGO_ENABLED=0

publish_ice = usr/publish/ice.$(shell go env GOOS).$(shell go env GOARCH)

all: def
	@echo && date
	go build -v -o bin/ice.bin src/main.go src/version.go src/binpack.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart
	rm ${publish_ice}; cp bin/ice.bin ${publish_ice}

def:
	@ [ -f src/version.go ] || echo "package main" > src/version.go
	@ [ -f src/binpack.go ] || echo "package main" > src/binpack.go

ice:
	cat src/binpack.go|sed 's/package main/package ice/g' > usr/release/binpack.go

app:
	CGO_ENABLED=1 go build -o usr/publish/contexts.app/Contents/MacOS/contexts src/app.go src/binpack.go && open usr/publish/contexts.app
	hdiutil create usr/publish/tmp.dmg -ov -volname "ContextsInstall" -fs HFS+ -srcfolder "usr/publish/contexts.app"
	rm -f usr/publish/ContextsInstall.dmg
	hdiutil convert usr/publish/tmp.dmg -format UDZO -o usr/publish/ContextsInstall.dmg

%: src/%.go
	@echo && date
	go build -v -o usr/publish/$@ src/$@.go && chmod u+x usr/publish/$@


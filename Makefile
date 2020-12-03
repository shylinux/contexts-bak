export GOPROXY=https://goproxy.cn
export GOPRIVATE=github.com
export CGO_ENABLED=0

all:
	@echo && date
	go build -v -o bin/ice.bin src/main.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart

relay: src/relay.go
	@echo && date
	go build -v -o usr/publish/$@ src/relay.go && chmod u+x usr/publish/$@

relay.darwin: src/relay.go
	@echo && date
	GOOS=darwin go build -v -o usr/publish/relay.darwin.amd64 src/relay.go

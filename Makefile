all:
	@echo && date
	GOPRIVATE=github.com GOPROXY=https://goproxy.cn CGO_ENABLED=0 go build -o bin/ice.bin src/main.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart

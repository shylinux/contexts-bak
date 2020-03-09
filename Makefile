all:
	@echo && date
	CGO_ENABLED=0 GOPRIVATE=github.com GOPROXY=https://goproxy.cn go build -o bin/ice.bin src/main.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart

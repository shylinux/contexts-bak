export GOPROXY=https://goproxy.cn
export GOPRIVATE=github.com
# export CGO_ENABLED=0

meta_info = "-X main.Time=`date +"%Y-%m-%d_%H:%M:%S"` -X main.Version=`git describe --tags` -X main.HostName=`hostname` -X main.UserName=`whoami`" 
all:
	@echo && date
	go build -ldflags $(meta_info) -v -o bin/ice.bin src/main.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart

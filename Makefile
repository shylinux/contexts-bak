
BENCH=src/examples/bench.go

install:
	@go get github.com/go-sql-driver/mysql
	@go get github.com/nsf/termbox-go
	@go get github.com/skip2/go-qrcode
	@go get github.com/gomarkdown/markdown
	go install $(BENCH)
	@date
	bench web.code.counter nmake 1

install_all: install
	touch etc/local.shy
	touch etc/local_exit.shy
	touch etc/init.shy
	touch etc/exit.shy
	touch etc/login.txt
	touch etc/history.txt

run:
	etc/bootstrap.sh

tar:
	[ -e tar ] || mkdir tar
	[ -e tar/bin ] || mkdir tar/bin
	[ -e tar/etc ] || mkdir tar/etc
	cp etc/bootstrap.sh tar/
	cp etc/init.shy tar/etc/
	cp etc/exit.shy tar/etc/
	touch tar/etc/local.shy
	touch tar/etc/exit_local.shy
	[ -e tar/usr ] || mkdir tar/usr
	cp -r usr/template tar/usr
	cp -r usr/librarys tar/usr
	[ -e tar/var ] || mkdir tar/var

tar_all: tar darwin linux64
	cp etc/local.shy tar/etc/
	cp etc/exit_local.shy tar/etc/
	mv bench.darwin tar/bin/
	mv bench.linux64 tar/bin/
	tar zcvf tar.tgz tar

linux64:
	GOARCH=amd64 GOOS=linux go build $(BENCH) -o bench.linux64
linux32:
	GOARCH=386 GOOS=linux go build $(BENCH) -o bench.linux32
linux_arm:
	GOARCH=arm GOOS=linux go build $(BENCH) -o bench.linux.arm
darwin:
	GOARCH=amd64 GOOS=darwin go build $(BENCH) -o bench.darwin

win64:
	GOARCH=amd64 GOOS=windows go build $(BENCH) -o bench.win64.exe
win32:
	GOARCH=386 GOOS=windows go build $(BENCH) -o bench.win32.exe


DOTS=etc/dotsfile
back_dotsfile:
	cp ~/.zshrc $(DOTS)
	cp ~/.tmux.conf $(DOTS)
	cp ~/context/.git/hooks/post-commit $(DOTS)/git_hooks
	cp ~/.vimrc $(DOTS)
	cp ~/.vim/syntax/shy.vim $(DOTS)

load_dotsfile:\
   	~/.zshrc\
   	~/.tmux.conf\
   	~/context/.git/hooks/post-commit\
   	~/.vimrc\
   	~/.vim/syntax/shy.vim

~/.zshrc: $(DOTS)/.zshrc
	cp $< $@
~/.tmux.conf: $(DOTS)/.tmux.conf
	cp $< $@
~/context/.git/hooks/post-commit: $(DOTS)/git_hooks/post-commit
	cp $< $@
~/.vimrc: $(DOTS)/.vimrc
	cp $< $@
~/.vim/syntax/shy.vim: $(DOTS)/shy.vim
	cp $< $@

.PHONY: tar run


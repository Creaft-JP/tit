all: install
tit:
	go build
install: tit
	mv tit /usr/local/bin/
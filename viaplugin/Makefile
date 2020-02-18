GOOS=$(shell go env GOOS)

all: mainapp goplugin cgoplugin

mainapp:
	go build -o ./bin/mainapp-${GOOS} ./

goplugin:
	go build -o ./bin/goplugin-${GOOS}.goos -buildmode=plugin ./plugin/goplugin

cgoplugin:
	go build -o ./bin/cgoplugin-${GOOS}.goos -buildmode=plugin ./plugin/cgoplugin

clean:
	rm -rf ./bin/*

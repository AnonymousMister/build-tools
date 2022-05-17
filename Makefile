build: build-liunx build-windows build-darwin

build-liunx:main.go
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-s -w" -a -o build-tools .
	mv build-tools ./lib/tools-liunx


build-windows:main.go
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -ldflags "-s -w" -a -o build-tools .
	mv build-tools ./lib/tools-windows

build-darwin:main.go
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -ldflags "-s -w" -a -o build-tools .
	mv build-tools ./lib/tools-darwin

push-linux: build-liunx
	scp lib/tools-liunx root@192.168.1.227:/home/config

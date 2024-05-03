.PHONY: run
run:
	@go run cmd/main.go

.PHONY: serve-wasm
serve:
	@./server

./PHONY: build-server
build-server:
	@go build -o ./server  ./wasm/server.go

./PHONY: build-server-win
build-server-win:
	@GOOS=windows GOARCH=amd64 go build -o ./server.exe  ./wasm/server.go
.PHONY: buildrun
buildrun:
	@ $(MAKE) build && ./metaballs 

.PHONY: build
build:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build


.PHONY: build-win
build-win:
#
	@sudo apt-get install gcc-mingw-w64
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o ./metaballs.exe ./cmd

.PHONY: build-win-fyne
build-win-fyne:
	@go install github.com/fyne-io/fyne-cross@latest
	@fyne-cross windows -arch=amd64

.PHONY: build-linux
build-linux:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./metaballs ./cmd


.PHONY: build-linux-fyne
build-linux-fyne:
	@go install github.com/fyne-io/fyne-cross@latest
	@fyne-cross linux -arch=amd64


.PHONY: build-wasm
build-wasm:
	@GOOS=js GOARCH=wasm go build -o ./wasm/metaballs.wasm ./cmd
	
.PHONY: pprof-cpu
pprof-cpu:
	@go tool pprof -http localhost:8080 profile/cpu.pprof

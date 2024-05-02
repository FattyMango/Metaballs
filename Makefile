.PHONY: run
run:
	@go run cmd/main.go

.PHONY: buildrun
buildrun:
	@ $(MAKE) build && ./metaballs 

.PHONY: build
build:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build


.PHONY: build-win
build-win:
#	@go install github.com/fyne-io/fyne-cross@latest
#	@fyne-cross windows -arch=amd64,386
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o ./metaballs.exe ./cmd


.PHONY: build-linux
build-linux:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./metaballs ./cmd
#	@go install github.com/fyne-io/fyne-cross@latest
#	@fyne-cross linux -
	

.PHONY: pprof-cpu
pprof-cpu:
	@go tool pprof -http localhost:8080 profile/cpu.pprof
.PHONY: run
run:
	@ $(MAKE) build && ./metaballs 

.PHONY: build
build:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build


.PHONY: build-win
build-win:
	@go install github.com/fyne-io/fyne-cross@latest
	@fyne-cross windows -arch=amd64,386
.PHONY: run
run:
	@ $(MAKE) build && ./metalballs

.PHONY: build
build:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build


.PHONY: build-win
build-win:
	@CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build 
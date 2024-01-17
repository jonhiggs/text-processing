.PHONY: build
build: build/unfold

build/%: cmd/%/main.go
	go build -o $@ $<

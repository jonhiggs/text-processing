.PHONY: build
build: build/unfold build/org build/demarkdown

build/%: cmd/$*/$*.go
	go build -o $@ $<

.PHONY: build
build: build/unfold build/org build/demarkdown

build/%: cmd/%/main.go
	go build -o $@ $<

clean:
	rm build/*

.PHONY: build

MANS := unfold.1
BINS := unfold org demarkdown

build: $(addprefix build/,$(BINS)) $(addprefix build/man/,$(MANS))

build/%: cmd/%/main.go
	go build -o $@ $<

build/man/%: export BUILD_DATE = $(shell date --iso-8601)
build/man/%:
	cat man/$* | envsubst '$${BUILD_DATE}' > $@

clean:
	rm -f build/* build/man/*

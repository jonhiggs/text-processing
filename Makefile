.PHONY: build

MANS := man1/unfold.1 man1/demarkdown.1 man1/org.1
BINS := unfold org demarkdown

build: \
	$(addprefix build/,$(BINS)) \
	$(addprefix build/man/,$(addsuffix .gz,$(MANS))) \
	$(MANS:%=doc/%.html)

build/%: cmd/%/main.go
	go mod tidy
	go test cmd/$*/*.go
	go build -o $@ $<

build/man/%.gz: export BUILD_DATE = $(shell date --iso-8601)
build/man/%.gz: | build/man/man1
	cat man/$* | envsubst '$${BUILD_DATE}' > build/man/$*
	gzip build/man/$*

doc/%.html: build/man/%.gz | doc/man1
	zcat $< | groff -mandoc -Thtml > $@

install: prefix ?= /usr/local
install:
	cp $(addprefix build/,$(BINS)) $(prefix)/bin
	cp $(MANS:%=build/man/%.gz) $(prefix)/share/man/man1

clean:
	rm -f $(addprefix build/,$(BINS)) $(addprefix build/man/,$(MANS)) $(addprefix build/man/,$(addsuffix .gz,$(MANS)))

build/man/man1 doc/man1:
	mkdir -p $@

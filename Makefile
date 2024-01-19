.PHONY: build

MANS := unfold.1 demarkdown.1 org.1
BINS := unfold org demarkdown

build: $(addprefix build/,$(BINS)) $(addprefix build/man/,$(addsuffix .gz,$(MANS)))

build/%: cmd/%/main.go
	go mod tidy
	go test cmd/$*/*.go
	go build -o $@ $<

build/man/%.gz: export BUILD_DATE = $(shell date --iso-8601)
build/man/%.gz:
	cat man/$* | envsubst '$${BUILD_DATE}' > build/man/$*
	gzip build/man/$*

clean:
	rm -f $(addprefix build/,$(BINS)) $(addprefix build/man/,$(MANS)) $(addprefix build/man/,$(addsuffix .gz,$(MANS)))

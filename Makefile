all: build/script2matrix

PREFIX = /usr

GO_BUILDFLAGS =
GO_LDFLAGS    =

build/script2matrix: FORCE
	go build $(GO_BUILDFLAGS) -ldflags '-s -w $(GO_LDFLAGS)' -o $@ .

check: all static-check build/cover.html FORCE
	@printf "\e[1;32m>> All tests successful.\e[0m\n"
static-check: FORCE
	@printf "\e[1;36m>> gofmt\e[0m\n"
	@if s="$$(gofmt -s -l *.go 2>/dev/null)" && test -n "$$s"; then printf ' => %s\n%s\n' gofmt  "$$s"; false; fi
	@printf "\e[1;36m>> go vet\e[0m\n"
	@go vet .
build/cover.out: FORCE
	@printf "\e[1;36m>> go test\e[0m\n"
	@go test -covermode count -coverprofile=$@ .
build/cover.html: build/cover.out
	go tool cover -html $< -o $@

install: FORCE all
	install -D -m 0755 build/script2matrix "$(DESTDIR)$(PREFIX)/bin/script2matrix"

vendor: FORCE
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: FORCE

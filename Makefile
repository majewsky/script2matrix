all: build/script2matrix

PKG    = github.com/majewsky/script2matrix
PREFIX = /usr

GO            = GOPATH=$(CURDIR)/.gopath GOBIN=$(CURDIR)/build go
GO_BUILDFLAGS =
GO_LDFLAGS    = -s -w

build/script2matrix: FORCE
	$(GO) install $(GO_BUILDFLAGS) -ldflags '$(GO_LDFLAGS)' '$(PKG)'

check: all static-check build/cover.html FORCE
	@printf "\e[1;32m>> All tests successful.\e[0m\n"
static-check: FORCE
	@if ! hash golint 2>/dev/null; then printf "\e[1;36m>> Installing golint...\e[0m\n"; go get -u golang.org/x/lint/golint; fi
	@printf "\e[1;36m>> gofmt\e[0m\n"
	@if s="$$(gofmt -s -l *.go 2>/dev/null)" && test -n "$$s"; then printf ' => %s\n%s\n' gofmt  "$$s"; false; fi
	@printf "\e[1;36m>> golint\e[0m\n"
	@if s="$$(golint . 2>&1)" && test -n "$$s"; then printf ' => %s\n%s\n' golint "$$s"; false; fi
	@printf "\e[1;36m>> go vet\e[0m\n"
	@$(GO) vet .
build/cover.out: FORCE
	@printf "\e[1;36m>> go test\e[0m\n"
	@$(GO) test -covermode count -coverprofile=$@ .
build/cover.html: build/cover.out
	$(GO) tool cover -html $< -o $@

install: FORCE all
	install -D -m 0755 build/script2matrix "$(DESTDIR)$(PREFIX)/bin/script2matrix"

.PHONY: FORCE

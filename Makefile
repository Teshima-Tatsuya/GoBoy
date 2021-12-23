NAME := GoBoy
BINDIR := ./build

.PHONY: build
build:
	@go build -tags macos -o $(BINDIR)/darwin-amd64/$(NAME)

.PHONY: clean
clean:
	@-rm -rf ./$(BINDIR)

.PNONY: test

test:
	make build
	$(BINDIR)/darwin-amd64/$(NAME) ./roms/hello.gb
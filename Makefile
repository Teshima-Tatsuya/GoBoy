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

test_cpu:
	go test ./pkg/gb/cpu

test_register:
	go test -run Register ./pkg/gb/cpu -cover -coverprofile=cover_register.out
test_opcode:
	go test -run OpCode ./pkg/gb/cpu -cover


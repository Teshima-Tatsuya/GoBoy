NAME := GoBoy
BINDIR := ./build

.PHONY: build
build:
	@go build -tags release -o $(BINDIR)/darwin-amd64/$(NAME) ./cmd/

build_debug:
	@go build -tags debug -o $(BINDIR)/darwin-amd64/$(NAME) ./cmd/

.PHONY: clean
clean:
	@-rm -rf ./$(BINDIR)

.PNONY: test

test_blargg:
	make build
	$(BINDIR)/darwin-amd64/$(NAME) ./test/blargg/cpu_instrs/cpu_instrs.gb

test_cpu:
	go test ./pkg/gb/cpu

test_register:
	go test -run Register ./pkg/gb/cpu -cover -coverprofile=cover_register.out
test_opcode:
	go test -run OpCode ./pkg/gb/cpu -cover


BINARY = backend
ifeq ($(OS),Windows_NT)
    BINARY = backend.exe
endif

dev: build
	./$(BINARY)

build: main.go
	go build -o ./$(BINARY) ./main.go

build-release: main.go
	go build -ldflags "-s -w" -o ./$(BINARY) ./main.go

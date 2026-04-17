dev: build
	./backend.exe

build: main.go
	go build -o ./backend.exe ./main.go

build-release: main.go
	go build -ldflags "-s -w" -o ./backend-release.exe ./main.go
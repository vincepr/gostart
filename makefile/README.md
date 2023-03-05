# makefile - build

## basics

```
// run it:
go run main.go

// build it out:
go build -o ./target/hello main.go

// run it in terminal
./target/hello
```

## build for different platforms:
- linux:
```
GOARCH=amd64 GOOS=linux go build -o ./target/hello-linux main.go
```

- windows:
```
GOARCH=amd64 GOOS=windows go build -o ./target/hello-win63.exe main.go
```

- windows 32bit:
```
GOARCH=386 GOOS=windows go build -o ./target/hello-win32 main.go
```

- or a linux arm-distro:
```
GOARCH=arm GOOS=linux go build -o ./target/hello-linux-arm main.go
```

## makefile
To order and structure support for different builds we can make a file `Makefile`
-  build with running: `make build` or `make run`
- or `make` for the whole file
- **IMPORTANT Note** golang-Makefile needs full Tabs.
- `@go build -o ./target/${BINDARY_NAME}-linux main.go` At-Symbol to remove console-output for that command

```Makefile
BINDARY_NAME = hello
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINDARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINDARY_NAME}-win64.exe main.go
	GOARCH=386 GOOS=windows go build -o ./target/${BINDARY_NAME}-win32 main.go
	GOARCH=arm GOOS=linux go build -o ./target/${BINDARY_NAME}-linux-arm main.go

run: build
	./target/${BINDARY_NAME}-linux

```

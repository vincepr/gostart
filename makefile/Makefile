BINDARY_NAME = hello
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINDARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINDARY_NAME}-win64.exe main.go
	GOARCH=386 GOOS=windows go build -o ./target/${BINDARY_NAME}-win32 main.go
	GOARCH=arm GOOS=linux go build -o ./target/${BINDARY_NAME}-linux-arm main.go

run: build
	./target/${BINDARY_NAME}-linux

clean:
	go clean
	rm ./target/${BINDARY_NAME}-linux main.go
	rm ./target/${BINDARY_NAME}-win64.exe main.go
	rm ./target/${BINDARY_NAME}-win32 main.go
	rm ./target/${BINDARY_NAME}-linux-arm main.go

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out
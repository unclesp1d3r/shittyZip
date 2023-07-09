BINARY_NAME=szip
BIN_DIR=bin

build:
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ${BIN_DIR}/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ${BIN_DIR}/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ${BIN_DIR}/${BINARY_NAME}-windows main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BIN_DIR}/${BINARY_NAME}-darwin
	rm ${BIN_DIR}/${BINARY_NAME}-linux
	rm ${BIN_DIR}/${BINARY_NAME}-windows
dep:
	go mod download
	go mod tidy
	go mod vendor

all: clean build

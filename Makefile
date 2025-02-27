BINARY_NAME=falco_postprocessor

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux-amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows-amd64 main.go
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}-linux-arm64 main.go
	GOARCH=arm64 GOOS=windows go build -o ${BINARY_NAME}-windows-arm64 main.go

clean:
	go clean
	rm ${BINARY_NAME}-darwin-amd64
	rm ${BINARY_NAME}-linux-amd64
	rm ${BINARY_NAME}-windows-amd64
	rm ${BINARY_NAME}-darwin-arm64
	rm ${BINARY_NAME}-linux-arm64
	rm ${BINARY_NAME}-windows-arm64

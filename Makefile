BINARY := matcher
test:
	echo "hello"
fmt:
	gofmt -w .
vet:
	go vet .
build: fmt vet
	go build -o $(BINARY) -tags=jsoniter
linux-build: fmt vet
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BINARY) -tags=jsoniter
windows-build: fmt vet
	env GOOS=windows CGO_ENABLED=0  go build -o $(BINARY).exe -tags=jsoniter
test: build
	./$(BINARY) test
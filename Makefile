# Parameters
MAIN_FILE=main.go
SERVICE=consumer

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o $(SERVICE)/$(SERVICE) $(SERVICE)/$(MAIN_FILE)

build-linux:
	GOARCH=amd64 GOOS=linux go build -o $(SERVICE)/$(SERVICE) $(SERVICE)/$(MAIN_FILE)

test:
	go test -race -v ./...
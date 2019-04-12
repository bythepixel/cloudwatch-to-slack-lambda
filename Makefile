M = $(shell printf "\033[34;1m▶\033[0m")

go-exists:
	@go version > /dev/null 2>&1 && echo $? || (echo "docker not found $$?"; exit 1)

build: go-exists build/main
	@GOOS=linux go build -o build/main main.go

pack: build
	@zip build/deployment.zip build/main

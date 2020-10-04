# This one is to automatically insert verison string
# from git tag output.
build:
	go build -ldflags="-X 'github.com/santosh/cotu/cmd.Version=`git describe --tags`'"

run:
	go run main.go

test:
	go test -v ./...


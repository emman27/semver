build:
	GOOS=darwin GOARCH=386 go build -o bin/semver-darwin_i386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/semver-x86_64 main.go

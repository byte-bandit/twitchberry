build:
	@env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/twitchberry cmd/twitchberry/main.go

dev:
	@go build -o bin/twitchberry.exe cmd/twitchberry/main.go
	@bin/twitchberry.exe

build-test:
	@env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/test cmd/tester/tester.go

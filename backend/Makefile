# Setup tools (install Air)
.PHONY: tools
tools:
	go install github.com/cosmtrek/air@latest

build:
	go build -o app

run:
	./app

# Run the application with Air for live reloading
.PHONY: run-dev
run-dev:
	air
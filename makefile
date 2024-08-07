.PHONY: dev
dev:
	go build -o ./tmp/app ./cmd/main.go && air

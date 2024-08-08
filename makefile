.PHONY: dev
dev:
	templ generate && go build -o ./tmp/app . && air
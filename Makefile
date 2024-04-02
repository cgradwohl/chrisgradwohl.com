build:
	@templ generate
	@docker compose up --build

run: build
	@./bin/server

dev:
	@command -v entr >/dev/null 2>&1 || { echo >&2 "entr not found. Install it here: https://github.com/eradman/entr or use 'make run' for manual building and executing."; exit 1; }
	@ls app/*.go app/page/*.templ app/component/*.templ | entr -r make run

test:
	@go test -v ./...

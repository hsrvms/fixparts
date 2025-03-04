.PHONY: dev
dev:
	templ generate --watch & docker compose up

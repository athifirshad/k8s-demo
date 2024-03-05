VERSION ?= latest

build:
	@docker build -t ghcr.io/athifirshad/contacts-backend:$(VERSION) .

push:
	@docker push ghcr.io/athifirshad/contacts-backend:$(VERSION)
dev: 
	@docker compose up
run:
	@go run ./cmd/api

psql:
	docker exec -it postgres psql -U root contacts

createdb:
	docker exec -it postgres createdb --user=root --owner=root contacts
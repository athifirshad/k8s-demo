VERSION ?= latest

build:
	@docker build -t ghcr.io/athifirshad/demo:$(VERSION) .

push:
	@docker push ghcr.io/athifirshad/demo:$(VERSION)
dev: 
	@docker compose up
run:
	@go run ./cmd/api

psql:
	docker exec -it postgres psql -U root contacts


kill:
	kubectl delete rollout contacts-backend-rollout
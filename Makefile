VERSION ?= latest

build:
	@docker build -t ghcr.io/athifirshad/demo:$(VERSION) .

push:
	@docker push ghcr.io/athifirshad/demo:$(VERSION)
dev: 
	@docker compose up
run:
	@go run ./cmd/api

kill:
	kubectl delete rollout api-rollout
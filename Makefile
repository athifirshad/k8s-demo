build:
	@docker compose up -build
dev: 
	@docker compose up
run:
	@go run ./cmd/api

stop:
	@docker compose down	
live:
	@air

asynq web:
	@docker run --rm --name asynqmon -p  8080:8080 hibiken/asynqmon --redis-addr=host.docker.internal:6379

mailer test:
	openssl s_client -starttls smtp -connect localhost:1025

psql:
	docker exec -it db psql -U root contacts

createdb:
	docker exec -it db createdb --user=root --owner=root contacts
	
.PHONY: up
up:
	@docker compose up -d

.PHONY: down
down:
	@docker compose down
	@docker system prune --all
	@sudo rm -rf .data/

.PHONY: connect_gemini_bot
connect_gemini_bot:
	@docker-compose exec cryptor-trends-gemini-bot sh

.PHONY: connect_coin_bot
connect_coin_bot:
	@docker-compose exec cryptor-trends-coin-bot sh

.PHONY: connect_api
connect_api:
	@docker-compose exec cryptor-trends-api bash

.PHONY: ps
ps:
	@docker compose ps

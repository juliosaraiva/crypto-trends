.PHONY: up
up:
	@docker compose up -d

.PHONY: up_dev
up_dev:
	@docker compose -f docker-compose.local.yaml up -d

.PHONY: down
down:
	@docker compose down

.PHONY: down_with_clean
down_with_clean:
	@docker compose down
	@docker system prune --all
	@sudo rm -rf .data/

.PHONY: connect_gemini_bot
connect_gemini_bot:
	@docker compose exec cryptor-trends-gemini-bot sh

.PHONY: connect_coin_bot
connect_coin_bot:
	@docker compose exec cryptor-trends-coin-bot sh

.PHONY: connect_api
connect_api:
	@docker compose exec cryptor-trends-api bash

.PHONY: ps
ps:
	@docker compose ps

.PHONY: docker_clean
docker_clean:
	@docker system prune --all

build:
	docker build --tag spotter .

run: build
	docker run --rm -it -p 8080:8080 spotter

compose:
	docker compose up --build --remove-orphans
	docker compose down -v

compose-dev:
	docker compose -f docker-compose-dev.yaml up --remove-orphans
	docker compose down -v

compose-debug:
	docker compose -f docker-compose-debug.yaml up --remove-orphans
	docker compose down -v
build:
	docker build --tag spotter .

run: build
	docker run --rm -it -p 8080:8080 spotter

compose:
	docker compose up --remove-orphans
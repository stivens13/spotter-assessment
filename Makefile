init: build create-network

create-network:
	docker network create backend-network

build:
	docker build -f services/spotter-api/Dockerfile --tag spotter .
	docker build -f services/youtube-api/Dockerfile --tag youtube .
	docker build -f services/etl/Dockerfile --tag etl .

run: build
	docker run --rm -it -p 8080:8080 spotter

compose: 
	docker compose -f docker-compose.yaml up --build --remove-orphans
	docker compose down -v # ensures reproducibility between runs






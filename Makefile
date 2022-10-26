build-prod:
	docker-compose --env-file ./.env.prod up --build -d
build-dev:
	docker-compose --env-file ./.env.dev up --build -d
destroy:
	docker-compose down
stop:
	docker-compose stop
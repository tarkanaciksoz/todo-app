build-prod:
	docker-compose --env-file ./.env.prod up --build -d
build-test:
	docker-compose --env-file ./.env.test up --build -d
build-dev:
	docker-compose --env-file ./.env.dev up --build -d backend && docker-compose --env-file ./.env.dev up --build -d frontend
destroy:
	docker-compose down
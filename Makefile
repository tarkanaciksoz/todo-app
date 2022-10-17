build-prod:
	docker-compose --env-file ./.env.prod up --build -d
build-test:
	docker-compose --env-file ./.env.test up --build -d
build-dev:
	docker-compose --env-file ./.env.dev up --build -d
destroy-prod:
	docker image rm {todo-list_backend,todo-list-backend,todo-list_frontend,todo-list-frontend} || true && docker rm {todo-app-frontend,todo-app-backend,todo-app-mysql} || true && docker-compose -f docker/prod/docker-compose.yaml down --volumes
destroy-test:
	docker image rm {todo-list_backend,todo-list-backend,todo-list_frontend,todo-list-frontend} || true && docker rm {todo-app-frontend,todo-app-backend,todo-app-mysql} || true && docker-compose -f docker/test/docker-compose.yaml down --volumes
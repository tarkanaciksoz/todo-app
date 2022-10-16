start:
	docker-compose up --detach
stop:
	docker-compose stop
build:
	docker-compose up --build -d
destroy:
	docker image rm {todo-list_backend,todo-list-backend,todo-list_frontend,todo-list-frontend} || true && docker rm {todo-app-frontend,todo-app-backend,todo-app-mysql} || true && docker-compose down --volumes
#!/bin/bash

echo "`/sbin/ip route|awk '/default/ { print $3 }'` todo-app.localhost" >> /etc/hosts

CGO_ENABLED=0 APP_ENV=$1 go test --tags $1 ./...
CGO_ENABLED=0 GOOS=linux go build --tags $1 -o main main.go
APP_ENV=$1 ./main
FROM golang:alpine
ARG ENV
WORKDIR /app

RUN apk update && \
    apk add --no-cache bash

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN chmod +x docker/wait-for.sh
RUN chmod +x docker/test-and-build.sh

EXPOSE 9090

CMD docker/./wait-for.sh mysql:3306 --timeout=30 -- APP_ENV=$1 ./main
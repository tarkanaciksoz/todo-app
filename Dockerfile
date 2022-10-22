FROM golang:alpine
ARG ENV
WORKDIR /app

RUN apk update && \
    apk add --no-cache bash

COPY go.mod go.sum ./
RUN go mod download
COPY . .

#RUN chmod +x docker/wait-for.sh
#RUN chmod +x docker/test-and-build.sh
RUN CGO_ENABLED=0 GOOS=linux go build --tags $ENV -o main main.go
EXPOSE 9090

#CMD docker/./wait-for.sh mysql:3306 --timeout=30 -- docker/./test-and-build.sh $ENV
CMD APP_ENV=$ENV ./main
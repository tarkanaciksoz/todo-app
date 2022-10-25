FROM golang:1.19-alpine
ARG ENV
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN APP_ENV=$ENV go build -o main main.go

EXPOSE 9090
CMD APP_ENV=$ENV ./main
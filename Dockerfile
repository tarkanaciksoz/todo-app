FROM golang:alpine
ARG ENV
WORKDIR /app

RUN apk update && \
    apk add --no-cache bash

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build --tags $ENV -o main main.go

FROM scratch
COPY --from=builder /app/main .
EXPOSE 9090
CMD ["./main"]
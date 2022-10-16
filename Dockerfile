FROM golang:alpine AS Builder
ARG ENV

RUN apk update && apk add bash
RUN apk add --no-cache bash

WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .

#RUN CGO_ENABLED=0 go test --tags $ENV ./...
RUN CGO_ENABLED=0 GOOS=linux go build --tags $ENV -o main main.go

FROM scratch
COPY --from=builder /app/main .
COPY ./.env .
CMD ["./main"]
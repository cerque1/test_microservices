FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/service_1 ./cmd/service_1

FROM alpine:latest 

WORKDIR /app

COPY /migrations /app/migrations
COPY --from=builder /app/service_1 /app/service_1

CMD [ "./service_1" ]
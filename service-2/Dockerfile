FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/service_2 ./cmd/service_2

FROM alpine:latest 

WORKDIR /app

COPY /docs /app/docs
COPY --from=builder /app/service_2 /app/service_2

CMD [ "./service_2" ]
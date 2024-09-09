# Build stage
FROM golang:1.20 AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .

# Run stage
FROM debian:stable-slim

WORKDIR /root/
COPY --from=build /app/main .

CMD ["./main"]

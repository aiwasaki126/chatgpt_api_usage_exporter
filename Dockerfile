FROM golang:1.22.1 as builder

WORKDIR /src
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest

WORKDIR /app
COPY --from=builder /src/main .

EXPOSE 8080

ENTRYPOINT [ "./main" ]


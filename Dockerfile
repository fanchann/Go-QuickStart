#this for testing!
FROM golang:latest as builder

WORKDIR /GoQuickStart
COPY . .

RUN go mod download
RUN CGO=ENABLED=0 GOOS=linux go build -o bin/go-start main.go

## testing
FROM golang:latest

WORKDIR /app
COPY --from=builder /GoQuickStart/bin/go-start /app/bin/
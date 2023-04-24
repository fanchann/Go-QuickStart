FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o go-start .
RUN chmod +x go-start
RUN mv go-start /bin/


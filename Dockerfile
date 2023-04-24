FROM golang:latest

ENV DB_DRIVER=mysql
ENV DB_AUTH_USERNAME=
ENV DB_AUTH_PASSWORD=root
ENV DB_NAME=information_schema
ENV DB_URL=0.tcp.ap.ngrok.io
ENV DB_PORT=11376

WORKDIR /app
COPY . .

RUN go build -o go-start .
RUN chmod +x go-start
RUN mv go-start /bin/


WORKDIR /tests
RUN go-start -pkg=fanchann/api
RUN go mod download
CMD [ "go","run","main.go" ]
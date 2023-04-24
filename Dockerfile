#this for testing!

FROM golang:latest


ENV DB_DRIVER=$DB_DRIVER
ENV DB_AUTH_USERNAME=$DB_AUTH_USERNAME
ENV DB_AUTH_PASSWORD=$DB_AUTH_PASSWORD
ENV DB_NAME=$DB_NAME
ENV DB_URL=$DB_URL
ENV DB_PORT=$DB_PORT

WORKDIR /app
COPY . .

RUN go build -o go-start .
RUN chmod +x go-start
RUN mv go-start /bin/


WORKDIR /tests
RUN go-start -pkg=fanchann/api
RUN go mod download
RUN apt update -y && apt install tree -y && tree >> dir.txt && cat dir.txt
CMD [ "go","run","main.go" ]
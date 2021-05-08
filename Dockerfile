FROM golang
WORKDIR /src/golang-api
ADD . /src/golang-api

RUN apt-get update 
RUN apt install -y git
RUN cd /src/golang-api 
RUN go get github.com/go-sql-driver/mysql 
RUN go get github.com/rs/cors 
RUN go build

EXPOSE 8080

ENTRYPOINT ./golang-api


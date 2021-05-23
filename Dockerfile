FROM golang
WORKDIR /go/src/golang-api
ADD . /go/src/golang-api

RUN apt-get update \
     && apt install -y git

RUN cd /go/src/golang-api \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/cors \
    && go build 

RUN mv golang-api /go/src/go-api \
    && cd /go/src \
    && rm -r golang-api

EXPOSE 8778

ENTRYPOINT /go/src/go-api


FROM golang
WORKDIR /src/golang-api
ADD . /src/golang-api

RUN apt-get update \
     && apt install -y git

RUN cd /src/golang-api \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/cors \
    && go build 
    
EXPOSE 8778

ENTRYPOINT ./golang-api


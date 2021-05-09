FROM zhung1027/golang-api
WORKDIR /go/src/golang-api1
ADD . /go/src/golang-api1

RUN apt-get update \
     && apt install -y git

RUN cd /go/src/golang-api1 \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/cors \
    && go build 
    
EXPOSE 8778

ENTRYPOINT ./golang-api


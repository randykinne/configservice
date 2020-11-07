FROM golang:alpine

RUN apk add git
RUN mkdir /app 

ADD . /app

RUN ls /app/src
WORKDIR /app/src
RUN go build -o main . 
CMD ["/app/src/main"]
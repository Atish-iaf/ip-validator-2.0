FROM golang:1.17-alpine

LABEL email="atish.iaf@gmail.com"

EXPOSE 8080

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main 

CMD [ "/app/main" ]
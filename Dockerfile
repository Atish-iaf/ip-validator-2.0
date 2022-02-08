FROM golang:1.17 AS builder
EXPOSE 8080
RUN mkdir /app
ADD . /app
WORKDIR /app
LABEL email="atish.iaf@gmail.com"
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./

FROM alpine:latest AS production
COPY --from=builder /app .
CMD [ "./main" ]
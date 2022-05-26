FROM golang:1.18

RUN mkdir /app
ADD . /app
WORKDIR /app

EXPOSE 8080

RUN go build -o apiserver cmd/main/main.go

CMD [ "./apiserver" ]

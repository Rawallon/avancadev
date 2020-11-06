FROM golang:1.15

WORKDIR /go/src/avancadev

COPY . .

RUN GOOS=linux go build 

EXPOSE 9090

ENTRYPOINT [ "./avancadev" ]
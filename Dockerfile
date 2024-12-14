FROM --platform=linux/amd64 docker.io/golang:alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOSUMDB=off

WORKDIR /app

COPY . .

COPY swagger/* /swagger/

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server ./main.go

WORKDIR /app

EXPOSE 8050

CMD ["/app/server"]
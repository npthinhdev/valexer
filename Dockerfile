FROM golang:latest
WORKDIR $GOPATH/src/valexer
COPY . .
EXPOSE 8080
CMD [ "go", "run", "server.go" ]
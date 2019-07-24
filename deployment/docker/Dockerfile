FROM golang:latest
WORKDIR $GOPATH/src/valexer
COPY valexer .
COPY web ./web
EXPOSE 8080
CMD ["./valexer"]
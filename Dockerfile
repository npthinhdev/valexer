FROM golang:latest
WORKDIR $GOPATH/src/github.com/npthinhdev/valexer
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only
RUN go build
EXPOSE 8080
CMD ["./valexer"]
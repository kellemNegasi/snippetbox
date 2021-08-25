FROM golang:1.16-alpine
RUN mkdir  /snippetbox
ADD . /snippetbox
WORKDIR /snippetbox
RUN go mod download
RUN go build -o main cmd/web/!(*_test).go
CMD ["/snippetbox/main"]

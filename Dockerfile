FROM golang:1.16-alpine
RUN mkdir  /snippetbox
ADD . /snippetbox
WORKDIR /snippetbox
RUN go mod download
RUN go build -o main cmd/web/*
EXPOSE 4000
CMD [".snippetbox/main"]

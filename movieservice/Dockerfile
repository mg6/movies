FROM golang:1.10

WORKDIR /go/src/github.com/mg6/movies
COPY . ./movieservice

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["movieservice"]

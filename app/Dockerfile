FROM golang:1.10.0

WORKDIR /go/src/app
COPY . .

RUN go get -d -v -t 
RUN go install -v

CMD ["app"]
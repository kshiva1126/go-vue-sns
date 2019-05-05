FROM golang:1.12.4

WORKDIR /go/src/go-vue-sns
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
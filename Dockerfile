FROM golang:1.13.8-alpine3.11

RUN apk add --update git

ENV GOBIN /go/bin

RUN go get github.com/tidwall/gjson
RUN go get github.com/tidwall/sjson
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq
RUN go get github.com/google/uuid

COPY *.go app/splitpay/
COPY view /go/src/view
COPY models /go/src/models
WORKDIR app/splitpay/

CMD go run main.go
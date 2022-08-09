FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN go build -o /go/bin/timer

FROM scratch
COPY --from=builder /go/bin/timer /go/bin/timer
ENTRYPOINT ["/go/bin/timer"]

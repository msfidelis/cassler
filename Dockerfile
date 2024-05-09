FROM golang:1.22 AS builder

WORKDIR $GOPATH/src/cassler

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cassler .


FROM scratch

COPY --from=builder /go/src/cassler/cassler ./

ENTRYPOINT ["./cassler"]
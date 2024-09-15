FROM golang:latest AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fingerd .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/app/fingerd .
COPY --from=builder /go/src/app/plans /root/plans/

EXPOSE 79

CMD ["./fingerd"]

FROM golang:1.12 AS builder
ENV GO111MODULE on
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /usr/bin/gdr .

###############################################

FROM ubuntu:18.04
COPY --from=builder /usr/bin/gdr /usr/bin/gdr

ENTRYPOINT ["/usr/bin/gdr"]

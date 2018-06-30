# ---- Build ----
FROM golang:1.10 as builder
WORKDIR /go/src/github.com/kdobmayer/hellows
COPY app/ .
RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo ./...

# ---- Final ----
FROM alpine:3.7
RUN apk --no-cache add ca-certificates curl jq
COPY --from=builder /go/bin/hellows /bin
COPY scripts/ready.sh /usr/bin

CMD ["hellows"]
EXPOSE 8080


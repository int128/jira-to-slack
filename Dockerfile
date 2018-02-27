FROM golang:1.10-alpine AS builder
WORKDIR /go/src/github.com/int128/jira-to-slack
RUN apk update && apk add --no-cache git
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
RUN apk update && apk add --no-cache ca-certificates
EXPOSE 3000
USER daemon
COPY --from=builder /go/bin/jira-to-slack /
CMD ["/jira-to-slack"]

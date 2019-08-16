FROM golang:1.12-alpine AS builder
RUN apk update && apk add --no-cache git gcc musl-dev
WORKDIR /build
COPY . .
RUN go install -v

FROM alpine
RUN apk update && apk add --no-cache ca-certificates
EXPOSE 3000
USER daemon
COPY --from=builder /go/bin/jira-to-slack /
CMD ["/jira-to-slack"]

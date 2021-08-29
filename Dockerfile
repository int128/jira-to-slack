FROM golang:1.16 as builder

WORKDIR /builder
COPY go.* .
RUN go mod download
COPY Makefile .
COPY main.go .
COPY pkg pkg
RUN make

FROM gcr.io/distroless/base-debian10
COPY --from=builder /builder/kubelogin /
EXPOSE 3000
CMD ["/jira-to-slack"]

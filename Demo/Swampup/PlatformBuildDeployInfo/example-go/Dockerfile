FROM golang:1.16-alpine AS builder

ARG srcpath="/build/jiracloud-connector"

RUN apk --no-cache add git && \
    mkdir -p "$srcpath"

ADD . "$srcpath"

RUN cd "$srcpath" && \
    GO111MODULE=on GOPRIVATE="${GOPRIVATE:-jfrog.com}" GONOPROXY="${GONOPROXY:-none}" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o /example-go

# Create Log folder
RUN mkdir -p /log

FROM gcr.io/distroless/base

COPY --from=builder /example-go /

CMD ["/example-go"]


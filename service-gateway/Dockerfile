##
## Build Stage
##
FROM golang:alpine AS builder

MAINTAINER 1ambda

ARG SRC_ROOT
ENV SRC_ROOT=${SRC_ROOT}

WORKDIR /go/src/${SRC_ROOT}/
COPY ./ /go/src/${SRC_ROOT}/
RUN pwd && ls -al ./ && ls -al ./internal
RUN ./.BUILD.sh

##
## Final Stage
##
FROM alpine
RUN apk --no-cache add ca-certificates

ARG SRC_ROOT
ENV SRC_ROOT=${SRC_ROOT}

COPY --from=builder /go/src/$SRC_ROOT/bin/server /app/

RUN addgroup -g 1000 -S go && \
    adduser -u 1000 -S go -G go
RUN chown -R go:go /app
USER go

ENTRYPOINT ["/app/server"]


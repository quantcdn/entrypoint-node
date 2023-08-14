FROM --platform=$BUILDPLATFORM golang:1.20 AS builder

ARG VERSION
ARG COMMIT

ADD . $GOPATH/src/github.com/quantcdn/entrypoint-node/

WORKDIR $GOPATH/src/github.com/quantcdn/entrypoint-node

ENV CGO_ENABLED 0

ARG TARGETOS TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} && \
    go mod tidy && \
    go build -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" -o build/entrypoint-node

FROM scratch

COPY --from=builder /go/src/github.com/quantcdn/entrypoint-node/build/entrypoint-node /usr/local/bin/entrypoint-node
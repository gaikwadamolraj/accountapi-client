FROM golang:1.17-alpine as builder

WORKDIR /go/src/github.com/gaikwadamolraj/form3/

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./

# Build Go Binary
RUN set -ex; \
    CGO_ENABLED=0 GOOS=linux go build ./*.go;

# Create final image
FROM alpine:latest

# Install Root Ceritifcates
RUN set -ex; \
    apk update; \
    apk add --no-cache \
     ca-certificates

WORKDIR /opt/

COPY --from=builder /go/src/github.com/gaikwadamolraj/form3/form3-client .

# Cretae appuser 
RUN addgroup -S appuser && adduser -S -G appuser appuser
RUN chown -R appuser:appuser /opt

USER appuser

# Run service with appuse
CMD /opt/form3-client
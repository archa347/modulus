FROM golang:1.14-alpine3.11 AS build
FROM alpine:3.11 AS base
# ================= BUILD ==================
FROM build AS builder
ARG PROJECT_BUILD_DIR=src/github.com/archa347/modulus
WORKDIR $GOPATH/$PROJECT_BUILD_DIR
COPY . .
RUN go build -o /build/modulus main.go

# ================== APP ===================
FROM base AS app
RUN addgroup -S app && adduser -S app -G app
USER app
WORKDIR app

COPY --from=builder --chown=app:app /build/modulus /app/modulus

ENTRYPOINT ["/app/modulus", "demo"]
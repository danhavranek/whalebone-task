# syntax=docker/dockerfile:1

# Build the app from source
# =========================

ARG GO_VERSION=1.24.4
FROM golang:${GO_VERSION} AS build-stage
# Set destination for COPY
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code
COPY . ./
# Install SQLite related packages
RUN apt-get update && apt-get install -y gcc libc6-dev libsqlite3-dev
# Build
RUN CGO_ENABLED=1 GOOS=linux go build -o /whalebone-task


# Run the tests in the container
# ==============================

# FROM build-stage AS run-test-stage
# RUN go test -v ./...


# Deploy the binary into a lean image
# ===================================

FROM gcr.io/distroless/base-debian12 AS build-release-stage
# Set destination for binary COPY
WORKDIR /
# Copy binary
COPY --from=build-stage /whalebone-task /whalebone-task
# Create DB volume
VOLUME ["/app/data"]

EXPOSE 8090
ENTRYPOINT ["/whalebone-task"]

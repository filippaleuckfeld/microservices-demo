# Start from the official Golang image with a specific version
FROM golang:1.20.4-alpine@sha256:0a03b591c358a0bb02e39b93c30e955358dadd18dc507087a3b7f3912c17fe13 AS builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base

# Set the working directory inside the container
WORKDIR /src

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /apiservice .

WORKDIR /src
COPY --from=builder /apiservice ./server

# Expose port 9090
EXPOSE 9090

# Set the entry point for the container
ENTRYPOINT ["/src/server"]

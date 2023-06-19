# # Start from the official Golang image with a specific version
# FROM golang:1.20.4

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go module files and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the application source code
# COPY . .

# # Build the Go application
# RUN go build -o main .

# # Expose port 9090
# EXPOSE 9090

# # Set the entry point for the container
# CMD ["./main"]

FROM golang:1.20.4-alpine@sha256:0a03b591c358a0bb02e39b93c30e955358dadd18dc507087a3b7f3912c17fe13 AS builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base

WORKDIR /src
# restore dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /apiservice .

FROM alpine:3.18.0@sha256:02bb6f428431fbc2809c5d1b41eab5a68350194fb508869a33cb1af4444c9b11
RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY --from=builder /apiservice ./server

ENV GOTRACEBACK=single

EXPOSE 9090
ENTRYPOINT ["/src/server"]

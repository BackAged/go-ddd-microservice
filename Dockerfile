# Defining App builder image
FROM golang:alpine AS builder

ARG VERSION=unversioned
ARG PRIVATE_KEY

# Add git to determine build git version
RUN apk add --no-cache --update git openssh

# Set GOPATH to build Go app
ENV GOPATH=/go

# Set apps source directory
ENV SRC_DIR=${GOPATH}/src/github.com/BackAged/go-ddd-microservice

# Define current working directory
WORKDIR ${SRC_DIR}

# Copy apps scource code to the image
COPY . ${SRC_DIR}

# Build App
RUN ./build.sh

# Adding the grpc_health_probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# Defining App image
FROM alpine:latest

RUN apk add --no-cache --update ca-certificates
RUN mkdir /app
# Copy App binary to image
COPY --from=builder /go/bin/order /usr/local/bin/order
COPY --from=builder /bin/grpc_health_probe /app/grpc_health_probe
EXPOSE 8000

ENTRYPOINT ["order"]
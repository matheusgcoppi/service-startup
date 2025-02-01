# Build The Go Binary.
FROM golang:1.23 as build_sales
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /service

# Build the services binary.
# WORKDIR is the working directory so you don't need to pass it in the go build command
WORKDIR /service/api/services/sales
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Run the Go Binary in Alpine.
# Runtime Stage: Lightweight Image with the Compiled Binary.
FROM alpine:3.21
ARG BUILD_DATE
ARG BUILD_REF
# Create a non-root user for security
RUN addgroup -g 1000 -S sales && \
    adduser -u 1000 -h /service -G sales -S sales
COPY --from=build_sales --chown=sales:sales /service/api/services/sales/sales /service/sales
WORKDIR /service
USER sales
CMD ["./sales"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="sales-api" \
      org.opencontainers.image.authors="Matheus Coppi" \
      org.opencontainers.image.source="https://github.com/matheusgcoppi/service" \
      org.opencontainers.image.revision="${BUILD_REF}"
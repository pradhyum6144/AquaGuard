# Build stage
FROM rust:1.74-slim AS builder

WORKDIR /app

# Copy manifests
COPY Cargo.toml Cargo.lock* ./
COPY src ./src

# Build the application
RUN cargo build --release

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

# Install SSL certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary from builder
COPY --from=builder /app/target/release/aquaguard-security .

# Expose port
EXPOSE 8081

# Run the application
CMD ["./aquaguard-security"]

FROM golang:alpine as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

RUN ls

# Copy and download dependency using go mod
RUN go mod download

RUN go get -t .

# Build the application
RUN go build -o server .

# Build final docker application to run
FROM alpine:3.9

# Move to /app directory as the place for resulting binary folder
WORKDIR /app

# Copy binary from build to main folder
COPY --from=builder /build/server .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/app/server"]
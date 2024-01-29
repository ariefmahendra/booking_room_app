# use official Golang image
FROM golang:1.21.6-alpine3.19 as build

# set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o booking-room

FROM alpine:3.19.0

WORKDIR /app
# Copy the binary from the build stage
COPY --from=build /app/booking-room .
COPY --from=build /app/.env .

# Run the executable
CMD ["./booking-room"]
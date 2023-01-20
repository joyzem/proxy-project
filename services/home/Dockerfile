FROM golang:alpine as builder

# Set the working directory to the app directory
WORKDIR /app/services/home

# Copy the source code into the container
COPY /services/home /app/services/home
COPY /services/utils /app/services/utils
COPY go.mod /app
COPY go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app .

WORKDIR /app/services/home

EXPOSE 80

CMD ["./main"]

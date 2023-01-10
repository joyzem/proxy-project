FROM golang:alpine as builder

# Set the working directory to the app directory
WORKDIR /app/services/account/frontend

# Copy the source code into the container
COPY /services/account/frontend /app/services/account/frontend
COPY .././go.mod /app
COPY .././go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app .

WORKDIR /app/services/account/frontend

EXPOSE 8083

CMD ["./main"]

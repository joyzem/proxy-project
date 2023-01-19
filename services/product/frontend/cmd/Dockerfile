FROM golang:alpine as builder

# Set the working directory to the app directory
WORKDIR /app/services/product/frontend/cmd

# Copy the source code into the container
COPY /services/product/frontend /app/services/product/frontend
COPY /services/product/domain /app/services/product/domain
COPY /services/utils /app/services/utils
COPY go.mod /app
COPY go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app .

WORKDIR /app/services/product/frontend/cmd

EXPOSE 8081

CMD ["./main"]

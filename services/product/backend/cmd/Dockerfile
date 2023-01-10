FROM golang:alpine as builder

# Set the working directory to the app directory
# The path must include "/services/product/backend" path because the code use s
# this path
WORKDIR /app/services/product/backend

# Copy the source code into the container
COPY /services/product/backend /app/services/product/backend
COPY /services/utils /app/services/utils
COPY ../../go.mod /app
COPY ../../go.sum /app

RUN go mod download
RUN cd /app/services/product/backend/cmd && go build -o main .

FROM alpine

# Set the working directory
WORKDIR /app/services/product/backend/cmd

# Copy the built binary from the builder stage
COPY --from=builder /app/services/product/backend/cmd/main .

EXPOSE 8080

CMD ["./main"]

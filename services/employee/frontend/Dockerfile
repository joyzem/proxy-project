FROM golang:alpine as builder

# Set the working directory to the app directory
WORKDIR /app/services/employee/frontend

# Copy the source code into the container
COPY /services/employee/frontend /app/services/employee/frontend
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app .

WORKDIR /app/services/employee/frontend

EXPOSE 8084

CMD ["./main"]

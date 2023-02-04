FROM golang:alpine as builder

# Set the working directory to the app directory
WORKDIR /app/services/organization/frontend

# Copy the source code into the container
COPY /services/organization/frontend /app/services/organization/frontend
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app .

WORKDIR /app/services/organization/frontend

EXPOSE 8082

CMD ["./main"]

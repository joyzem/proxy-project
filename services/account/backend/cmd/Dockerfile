FROM golang:alpine as builder

# Set the working directory to the app directory
# The path must include "/services/product/backend" path because the code use s
# this path
WORKDIR /app/services/account/backend/cmd

# Copy the source code into the container
COPY /services/account/backend /app/services/account/backend
COPY /services/account/dto /app/services/account/dto
COPY /services/account/domain /app/services/account/domain
COPY /services/base /app/services/base
COPY go.mod /app
COPY go.sum /app

RUN go build -o main .

FROM alpine

# Set the working directory
WORKDIR /app/services/account/backend/cmd

# Copy the built binary from the builder stage
COPY --from=builder /app/services/account/backend/cmd/main .

EXPOSE 7073

CMD ["./main"]

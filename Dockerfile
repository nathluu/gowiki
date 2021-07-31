##
## Build
##
FROM golang:1.16-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /gowiki

##
## Deploy
##
FROM alpine:3.14.0

WORKDIR /

COPY --from=build /gowiki /gowiki

EXPOSE 8080
ENTRYPOINT ["/gowiki"]
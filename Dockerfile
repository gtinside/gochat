# Dockerfile for generating the image for gochat application
FROM golang:latest
LABEL maintainer="Gaurav Tiwari <gtinside@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8090
CMD ["./chatserver"]
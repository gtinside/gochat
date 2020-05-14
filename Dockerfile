# Dockerfile for generating the image for gochat application
FROM golang:latest
LABEL maintainer="Gaurav Tiwari <gtinside@gmail.com>"
WORKDIR $GOPATH/src/github.com/gtinside/gochat
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8090
CMD ["chatserver"]
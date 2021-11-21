# Pull official ubuntu base image
FROM golang:latest

# Set working directory
WORKDIR /app

# Copy repo
COPY . /app/

RUN go get ./...

RUN go build ./...

CMD ./breadboi.tv

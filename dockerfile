#Base Image
FROM golang:1.23.1

#Working Directory
WORKDIR /app

#Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

#Copy the code
COPY . .

#Build App
RUN go build -o fetch-golang-api ./cmd/fetch-golang-api

ENTRYPOINT [ "/app/fetch-golang-api" ]
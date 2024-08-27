FROM golang:1.22.2
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build cmd/main.go
CMD ["/rest_server"]
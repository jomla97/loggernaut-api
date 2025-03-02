FROM golang:1.24-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o loggernaut-api
EXPOSE 80
CMD ["./loggernaut-api"]
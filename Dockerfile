FROM golang:1.22-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/server

EXPOSE 8080
CMD ["./server"]


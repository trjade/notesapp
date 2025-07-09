FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY .env .env

RUN go build -o notesapp ./cmd/app

CMD [ "./notesapp" ]

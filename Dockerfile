FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o notesapp ./cmd/app

EXPOSE 8080

CMD [ "./notesapp" ]

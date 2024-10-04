FROM golang:alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-libs sqlite-dev
ENV CGO_ENABLED=1

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./app

CMD ["./main"]
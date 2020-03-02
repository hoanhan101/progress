FROM golang:1.13

WORKDIR /app

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o bin/progress cmd/progress.go

EXPOSE 8000

ENTRYPOINT ["./bin/progress"]

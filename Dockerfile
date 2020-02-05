FROM golang:1.13

COPY . .
EXPOSE 8000

ENTRYPOINT ["./progress"]

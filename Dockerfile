FROM golang:1.13

COPY progress ./progress
EXPOSE 8000

ENTRYPOINT ["./progress"]

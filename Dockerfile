FROM golang:1.15 as builder
WORKDIR /app

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
ENV GODEBUG=madvdontneed=1
COPY --from=builder /app/app .
CMD ["./app"]

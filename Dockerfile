FROM golang:1.15

COPY main.go .

ENV GODEBUG=madvdontneed=1

CMD ["go", "run", "main.go"]

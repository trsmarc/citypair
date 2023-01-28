FROM golang:1.19 as builder

WORKDIR /api
COPY . .

WORKDIR /api/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server

FROM alpine:3.17.1

WORKDIR /app

COPY --from=builder /api/Makefile  /app/Makefile
COPY --from=builder /api/cmd/bin/server /app/bin/server
COPY --from=builder /api/config/*.yaml /app/config/
COPY --from=builder /api/cmd/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/bin/server
RUN chmod +x /app/entrypoint.sh

CMD ["./entrypoint.sh"]

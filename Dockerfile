FROM golang:1.23-alpine3.20 AS builder

COPY . .

RUN CGO_ENABLE=0 go build -o /uuidgen ./

FROM alpine:3.20 AS dist

COPY --from=builder /uuidgen /usr/local/bin/uuidgen

ENTRYPOINT ["uuidgen"]
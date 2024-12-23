FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o noop .

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/rjocoleman/noop
COPY --from=builder /app/noop /
ENTRYPOINT ["/noop"]
CMD ["/noop"]

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o redditclone .

# Minimal runtime (go images too big for this app)
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/redditclone .
COPY --from=builder /app/static ./static

EXPOSE 4040

CMD [ "./redditclone" ]
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o redditclone .

# Minimal runtime (go images too big for this app)
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app .

EXPOSE 4040

CMD [ "./redditclone" ]
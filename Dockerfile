
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o redditclone .

# Minimal runtime (go images too big for this app)
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/redditclone .

EXPOSE 25567

CMD [ "./redditclone" ]
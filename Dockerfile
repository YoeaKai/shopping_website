FROM golang:alpine as builder
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

FROM chromedp/headless-shell:latest
RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY . .
EXPOSE 8081
CMD ["/server/crawl_server"]
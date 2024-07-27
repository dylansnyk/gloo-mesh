FROM alpine:3.19.3

# installs public root certs
# needed to establish trust with third party sources
RUN apk upgrade --update-cache \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

COPY cert-agent-linux-amd64 /usr/local/bin/cert-agent

USER 10101

ENTRYPOINT ["/usr/local/bin/cert-agent"]

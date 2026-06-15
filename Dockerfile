FROM golang:1.22-alpine AS builder

ARG GITHUB_OWNER
ARG PNAPCTL_VERSION
RUN apk add --no-cache curl tar

WORKDIR /tmp

RUN curl -L -o pnapctl.tar.gz \
  https://github.com/${GITHUB_OWNER}/pnapctl/releases/download/${PNAPCTL_VERSION}/pnapctl-linux-amd64.tar.gz \
  && tar xzf pnapctl.tar.gz \
  && chmod +x pnapctl

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates bash \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /tmp/pnapctl /usr/local/bin/pnapctl

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

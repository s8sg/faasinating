FROM golang:1.9.2-alpine as builder

MAINTAINER alex@openfaas.com
ENTRYPOINT []

RUN env

RUN apk --no-cache add make curl git \
    && curl -sL https://github.com/openfaas/faas/releases/download/0.7.0/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog

WORKDIR /go/src/github.com/openfaas/faas/sample-functions/GoogleTranslator

COPY . /go/src/github.com/openfaas/faas/sample-functions/GoogleTranslator

RUN make install

FROM alpine:3.6

ARG test

# Needed to reach the hub
RUN apk --no-cache add ca-certificates 

COPY --from=builder /usr/bin/fwatchdog  /usr/bin/fwatchdog
COPY --from=builder /go/src/github.com/openfaas/faas/sample-functions/GoogleTranslator/googletranslator  /usr/bin/googletranslator

ENV fprocess "/usr/bin/googletranslator"
ENV TEST $test

HEALTHCHECK --interval=5s CMD [ -e /tmp/.lock ] || exit 1
CMD ["/usr/bin/fwatchdog"]

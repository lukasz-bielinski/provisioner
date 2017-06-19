FROM alpine:3.6

ENV KUBECTL_VERSION=v1.4.6


RUN apk add --no-cache ca-certificates 
RUN apk add --no-cache --virtual .build-deps \
  curl


RUN curl -L https://storage.googleapis.com/kubernetes-release/release/$KUBECTL_VERSION/bin/linux/amd64/kubectl -o /usr/bin/kubectl \
    && chmod +x /usr/bin/kubectl

RUN apk del .build-deps
COPY /src/bin/test /
ENTRYPOINT ["/test"]

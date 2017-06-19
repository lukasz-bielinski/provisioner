FROM alpine:3.6
RUN apk add --no-cache ca-certificates apache2-utils

COPY /src/bin/test /
ENTRYPOINT ["/test"]

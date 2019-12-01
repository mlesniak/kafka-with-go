# ------------------------------------------------------------------
FROM golang:alpine
# Update to edge for latest version of library
RUN sed -i -e 's/v[[:digit:]]\..*\//edge\//g' /etc/apk/repositories
RUN apk upgrade --update-cache --available
RUN apk add alpine-sdk git librdkafka librdkafka-dev pkgconfig
RUN go get -u github.com/mlesniak/kafka-with-go
RUN apk del librdkafka-dev pkgconfig alpine-sdk

# ------------------------------------------------------------------
FROM alpine:latest
# Add depdendent library, too. We tried to create a fully static binary, but ... failed (for now!).
COPY --from=0 /usr/lib/ /usr/lib
COPY --from=0 /usr/lib/librdkafka++.so.1 /usr/lib/librdkafka++.so.1

COPY --from=0 /go/bin/kafka-with-go /usr/local/bin
CMD ["/usr/local/bin/kafka-with-go"]